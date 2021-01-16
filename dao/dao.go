package dao

import "github.com/manicar2093/YoFioExamen/entities"

type CreditDetailsDao interface {
	GetAllCreditDetails() ([]entities.CreditDetails, error)
}
