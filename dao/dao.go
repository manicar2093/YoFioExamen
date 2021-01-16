package dao

import "github.com/manicar2093/YoFioExamen/entities"

type CreditDetailsDao interface {
	// GetAllCreditDetails obtiene todos los entities.CreditDetails que se deben validar
	GetAllCreditDetails() ([]entities.CreditDetails, error)
}
// CreditDetailsDaoImpl existe considerando que puede extraerse esta información de una base de datos.
// Solo se debe inyectar la dependencia necesaria de conexión e implementar la lógica
type CreditDetailsDaoImpl struct {}

// NewCreditDetailsDaoImpl crea una nueva instancia de CreditDetailsDaoImpl
func NewCreditDetailsDaoImpl() CreditDetailsDao{
	return &CreditDetailsDaoImpl{}
}

func (c CreditDetailsDaoImpl) GetAllCreditDetails() ([]entities.CreditDetails, error) {
	return []entities.CreditDetails{
		{LoanQuantity:300, Count: 0},
		{LoanQuantity:500, Count: 0},
		{LoanQuantity:700, Count: 0},
	}, nil

}
