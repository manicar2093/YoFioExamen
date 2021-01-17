package dao

import (
	"context"
	"github.com/manicar2093/YoFioExamen/entities"
	"github.com/manicar2093/YoFioExamen/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CreditDetailsDao interface {
	Save(data *entities.CreditDetailsWithStatus) error
	FilterCreditDetailsWithStatus(filter interface{}) ([]entities.CreditDetailsWithStatus, error)
}

type CreditDetailsDaoImpl struct {
	collection *mongo.Collection
}

func NewCreditDetailsDao(collection *mongo.Collection) CreditDetailsDao {
	return &CreditDetailsDaoImpl{collection: collection}
}

func (c CreditDetailsDaoImpl) Save(data *entities.CreditDetailsWithStatus) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	_, e := c.collection.InsertOne(ctx, data)
	if e != nil {
		return e
	}
	return nil
}

func (c CreditDetailsDaoImpl) FilterCreditDetailsWithStatus(filter interface{}) ([]entities.CreditDetailsWithStatus, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	var response []entities.CreditDetailsWithStatus
	cursor, e := c.collection.Find(ctx, filter)
	if e != nil {
		utils.LogError.Printf("Error al realizar el filtro de CreditDetails. \n\tDetalles: %v", e)
		return []entities.CreditDetailsWithStatus{}, e
	}
	if e := cursor.All(ctx, &response); e != nil {
		utils.LogError.Printf("Error extraer datos del cursor. \n\tDetalles: %v", e)
		return []entities.CreditDetailsWithStatus{}, e
	}
	return response, nil
}

