package services

import (
	"github.com/manicar2093/YoFioExamen/dao"
	"github.com/manicar2093/YoFioExamen/entities"
)

type CreditDetailsService interface {
	// GetAllCreditDetails obtiene todos los entities.CreditDetails que se deben validar
	GetAllCreditDetails() ([]entities.CreditDetails, error)
	SaveSuccessfulRequest(credit1 *entities.CreditDetails,credit2 *entities.CreditDetails,credit3 *entities.CreditDetails, invest int32) error
	SaveUnsuccessfulRequest(credit1 *entities.CreditDetails,credit2 *entities.CreditDetails,credit3 *entities.CreditDetails, invest int32) error
}

type CreditDetailsServiceImpl struct {
	creditDetailsDao dao.CreditDetailsDao
}

func NewCreditDetailsService(creditDetailsDao dao.CreditDetailsDao) CreditDetailsService {
	return &CreditDetailsServiceImpl{creditDetailsDao: creditDetailsDao}
}

// GetAllCreditDetails existe considerando que puede extraerse esta información de una base de datos.
// Solo se debe implementar la lógica necesaria
func (c CreditDetailsServiceImpl) GetAllCreditDetails() ([]entities.CreditDetails, error) {
	return []entities.CreditDetails{
		{LoanQuantity: 300, Count: 0},
		{LoanQuantity: 500, Count: 0},
		{LoanQuantity: 700, Count: 0},
	}, nil
}

func (c CreditDetailsServiceImpl) SaveSuccessfulRequest(credit1 *entities.CreditDetails, credit2 *entities.CreditDetails, credit3 *entities.CreditDetails, invest int32) error {
	panic("implement me")
}

func (c CreditDetailsServiceImpl) SaveUnsuccessfulRequest(credit1 *entities.CreditDetails, credit2 *entities.CreditDetails, credit3 *entities.CreditDetails, invest int32) error {
	panic("implement me")
}


