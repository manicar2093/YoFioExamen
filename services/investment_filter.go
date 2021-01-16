package services

import (
	"errors"
	"github.com/manicar2093/YoFioExamen/entities"
	"math"
)

// InvestmentFilter es una interfaz que se usará para realizar la lógica completa de filtrado
type InvestmentFilter interface {
	// Filter es la acción que se realizará para el calculo de creditos
	Filter(quantity int32, credit1, credit2, credit3 *entities.CreditDetails) (e error)
}

type InvestmentFilterImpl struct {}

func NewInvestmentFilter() InvestmentFilter {
	return &InvestmentFilterImpl{}
}

func (i InvestmentFilterImpl) Filter(quantity int32, credit1, credit2, credit3 *entities.CreditDetails) (e error) {
	var exponente, avoid int32 = 1, 0
	var creditsDetails []*entities.CreditDetails

	creditsDetails = append(creditsDetails, credit1)
	creditsDetails = append(creditsDetails, credit2)
	creditsDetails = append(creditsDetails, credit3)


	var loopCounter = func(amount int32) {
		for {
			if e = isLessThanZero(amount); isZero(amount) || e != nil {
				break
			}
			if can, divisor := isDivisible(amount,avoid, creditsDetails...); can {
				amount = amount - divisor.LoanQuantity
				avoid = divisor.LoanQuantity
				divisor.Count += exponente
			} else {
				avoid = 0
			}
		}
	}

	if can, _ := isDivisible(quantity,0, creditsDetails...); can {
		loopCounter(quantity)
		return e
	}

	q1, q2 := getThousandsAndRemaining(quantity)

	can, _ := isDivisible(q2,0, creditsDetails...)

	if !can {
		return NoCreditAssigment{Investment: quantity, Remaining: q2}
	}

	loopCounter(q1)

	loopCounter(q2)

	return e
}

// isDivisible valida si el quantity es divisible entre alguna de las LoanQuantity proporcionadas
// en los creditsDetails
func isDivisible(quantity int32, avoid int32, creditsDetails ...*entities.CreditDetails) (bool, *entities.CreditDetails) {
	for _, v := range creditsDetails {
		if quantity%v.LoanQuantity == 0 && v.LoanQuantity != avoid{
			return true, v
		}
	}
	return false, &entities.CreditDetails{}
}

// getThousandsAndRemaining disecciona la inversión para obtener los miles y el remanente en centenas
func getThousandsAndRemaining(quantity int32) (int32, int32) {
	// Dividimos la inversion entre 1000
	divided := quantity / 1000
	// Obtenemos la cantidad de miles que llegaron en la inversión
	floored := int32(math.Floor(float64(divided)))
	// Resolvemos la cantidad a miles
	thousands := floored * 1000

	return thousands, quantity - thousands

}

func isLessThanZero(q int32) error {
	if q < 0 {
		return errors.New("Error al realizar asignación. El calculo dió como resultado un numero negativo")
	}
	return nil
}

func isZero(q int32) bool {
	return q == 0
}

