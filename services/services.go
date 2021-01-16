package services

import (
	"fmt"
	"github.com/manicar2093/YoFioExamen/dao"
)

// NoCreditAssigment indica que la inversión no se puede asignar a los creditos establecidos
type NoCreditAssigment struct {
	Investment int32
	Remaining  int32
}

func (n NoCreditAssigment) Error() string {
	return fmt.Sprintf("Se tiene un remanente de $%d al asignar la inversión de $%d", n.Remaining, n.Investment)
}

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type CreditAssignerImpl struct {
	filter InvestmentFilter
	creditDetailsDao dao.CreditDetailsDao
}

func NewCreditAssigner(filter InvestmentFilter, creditDetailsDao dao.CreditDetailsDao) CreditAssigner {
	return &CreditAssignerImpl{
		filter: filter,
		creditDetailsDao: creditDetailsDao,
	}
}

func (c CreditAssignerImpl) Assign(investment int32) (int32, int32, int32, error) {
	cd, e := c.creditDetailsDao.GetAllCreditDetails()
	if e != nil {
		return 0, 0, 0, e
	}
	credit1 := &cd[0]
	credit2 := &cd[1]
	credit3 := &cd[2]

	e = c.filter.Filter(investment, credit1,credit2,credit3)
	if e != nil {
		return credit1.Count, credit2.Count, credit3.Count, e
	}
	return credit1.Count, credit2.Count, credit3.Count, nil
}

