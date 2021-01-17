package dao

import (
	"context"
	"github.com/manicar2093/YoFioExamen/connections"
	"github.com/manicar2093/YoFioExamen/entities"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestSave(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	data := entities.CreditDetailsWithStatus{
		Investment:     30000,
		CreditsDetails: []*entities.CreditDetails{
			{
				Count:        8,
				LoanQuantity: 300,
			},
			{
				Count:        8,
				LoanQuantity: 500,
			},
			{
				Count:        8,
				LoanQuantity: 500,
			},
		},
		Status:         "successful",
	}

	db := connections.GetMongoConnection(ctx,cancel)
	collection := db.Collection("credit_details")
	dao := NewCreditDetailsDao(collection)

	e := dao.Save(&data)
	assert.Nil(t, e, "No debió haber error al guardar el archivo")

	e = collection.Drop(context.Background())
	assert.Nil(t, e, "No debió haber error al borrar la base de datos")
}

func TestFilterCreditDetailsWithStatus(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	toStore := []entities.CreditDetailsWithStatus{
		{
			Investment:     30000,
			CreditsDetails: []*entities.CreditDetails{
				{
					Count:        8,
					LoanQuantity: 300,
				},
				{
					Count:        8,
					LoanQuantity: 500,
				},
				{
					Count:        8,
					LoanQuantity: 500,
				},
			},
			Status:         "successful",
		},
		{
			Investment:     15000,
			CreditsDetails: []*entities.CreditDetails{
				{
					Count:        8,
					LoanQuantity: 300,
				},
				{
					Count:        8,
					LoanQuantity: 500,
				},
				{
					Count:        8,
					LoanQuantity: 500,
				},
			},
			Status:         "successful",
		},
	}

	db := connections.GetMongoConnection(ctx,cancel)
	collection := db.Collection("credit_details")
	dao := NewCreditDetailsDao(collection)

	for _,v := range toStore {
		e := dao.Save(&v)
		assert.Nil(t, e, "No debió haber error al guardar el archivo")
	}

	data, e := dao.FilterCreditDetailsWithStatus(primitive.D{
		primitive.E{
			Key:   "status",
			Value: "successful",
		},
	})

	assert.Equal(t, len(toStore), len(data), "No se recogió la cantidad de datos esperada")
	
	e = collection.Drop(context.Background())
	assert.Nil(t, e, "No debió haber error al borrar la base de datos")
}