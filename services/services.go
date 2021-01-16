package services

import (
	"fmt"
	"github.com/manicar2093/YoFioExamen/dao"
	"github.com/manicar2093/YoFioExamen/entities"
)

// NoCreditAssigment indica que la inversión no se puede asignar a los creditos establecidos
type NoCreditAssigment struct {
	Investment int32
	Remaining  int32
}

func (n NoCreditAssigment) Error() string {
	return fmt.Sprintf("Se tiene un remanente de $%d al asignar la inversión de $%d", n.Remaining, n.Investment)
}

// InvestmentFilter es una interfaz que se usará para realizar la lógica completa de filtrado
type InvestmentFilter interface {
	// Filter es la acción que se realizará para el calculo de creditos
	Filter(quantity int32, credit1, credit2, credit3 *entities.CreditDetails) (e error)
}

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type CreditAssignerImpl struct {
	filter InvestmentFilter
	creditDetailsDao dao.CreditDetailsDao
}

func NewCreditAssigner(filter InvestmentFilter, creditDetailsDao dao.CreditDetailsDao) *CreditAssignerImpl {
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

