package dao

import (
	"github.com/manicar2093/YoFioExamen/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreditDetailsDao interface {
	FilterCreditDetails(filter interface{}) ([]entities.CreditDetails, error)
	FindStatistics() (entities.CreditsAssignmentStatistics, error)
}

type CreditDetailsDaoImpl struct {
	collections mongo.Collection
}

// NewCreditDetailsDaoImpl crea una nueva instancia de CreditDetailsDaoImpl
/*func NewCreditDetailsDaoImpl(collection mongo.Collection) CreditDetailsDao {
	return &CreditDetailsDaoImpl{collections: collection}
}

func (c CreditDetailsDaoImpl) FilterCreditDetails(filter interface{}) ([]entities.CreditDetails, error) {
	return []entities.CreditDetails{
		{LoanQuantity: 300, Count: 0},
		{LoanQuantity: 500, Count: 0},
		{LoanQuantity: 700, Count: 0},
	}, nil

}
*/