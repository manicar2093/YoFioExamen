package mocks

import (
	"github.com/manicar2093/YoFioExamen/entities"
	"github.com/stretchr/testify/mock"
)

// Services Mocks
type CreditServiceMock struct {
	mock.Mock
}

func (c CreditServiceMock) Assign(investment int32) (int32, int32, int32, error) {
	args := c.Called(investment)
	return args.Get(0).(int32), args.Get(1).(int32), args.Get(2).(int32), args.Error(3)

}

type InvestmentFilterMock struct {
	mock.Mock
}

func (i InvestmentFilterMock) Filter(quantity int32, credit1, credit2, credit3 *entities.CreditDetails) (e error) {
	args := i.Called(quantity, credit1, credit2, credit3)
	return args.Error(0)
}

// DAO Mocks

type CreditDetailsDaoMock struct {
	mock.Mock
}

func (c CreditDetailsDaoMock) GetAllCreditDetails() ([]entities.CreditDetails, error) {
	args := c.Called()
	return args.Get(0).([]entities.CreditDetails), args.Error(1)
}


