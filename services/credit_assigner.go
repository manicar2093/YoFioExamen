package services

import (
	"fmt"
	"github.com/manicar2093/YoFioExamen/utils"
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
	filter               InvestmentFilter
	creditDetailsService CreditDetailsService
}

func NewCreditAssigner(filter InvestmentFilter, creditAssignmentStatisticsService CreditDetailsService) CreditAssigner {
	return &CreditAssignerImpl{
		filter:               filter,
		creditDetailsService: creditAssignmentStatisticsService,
	}
}

func (c CreditAssignerImpl) Assign(investment int32) (int32, int32, int32, error) {
	cd, e := c.creditDetailsService.GetAllCreditDetails()
	if e != nil {
		return 0, 0, 0, e
	}
	credit1 := &cd[0]
	credit2 := &cd[1]
	credit3 := &cd[2]

	e = c.filter.Filter(investment, credit1, credit2, credit3)
	if e != nil {

		if noCreditAssigmentError, ok := e.(NoCreditAssigment); ok {
			e = c.creditDetailsService.SaveUnsuccessfulRequest(credit1, credit2, credit3, investment)
			if e != nil {
				utils.LogError.Printf("Sucedió un error al guardar los datos del calculo de la inversion de $%d.00 - Procesamiento NO EXITOSO. No se envió la información.", investment)
				return 0, 0, 0, e
			}
			return 0, 0, 0,noCreditAssigmentError
		}

		return 0, 0, 0, e
	}

	e = c.creditDetailsService.SaveSuccessfulRequest(credit1, credit2, credit3, investment)
	if e != nil {
		utils.LogError.Printf("Sucedió un error al guardar los datos del calculo de la inversion de $%d.00 - Procesamiento EXITOSO. No se envió la información.", investment)
		return 0, 0, 0, e
	}

	return credit1.Count, credit2.Count, credit3.Count, nil
}
