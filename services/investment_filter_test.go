package services

import (
	"fmt"
	"github.com/manicar2093/YoFioExamen/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sumAllCredits = func(credits ...*entities.CreditDetails) int32 {
	total := int32(0)
	for _, v := range credits {
		total += v.TotalWithCount()
	}
	return total
}

// TestInvestmentFilter valida el flujo sin errores
func TestInvestmentFilter(t *testing.T) {
	// Test data
	invest := int32(1000)
	credit1 := entities.CreditDetails{
		LoanQuantity: 300,
	}
	credit2 := entities.CreditDetails{
		LoanQuantity: 500,
	}
	credit3 := entities.CreditDetails{
		LoanQuantity: 700,
	}
	// Running Test
	service := NewInvestmentFilter()
	e := service.Filter(invest, &credit3, &credit1, &credit2)
	total := sumAllCredits(&credit1, &credit2, &credit3)

	// Test Validations
	assert.Nil(t, e, "No debió regresar error")
	assert.Equal(t, invest, total, "Las cantidades no coinciden!")
}

// TestInvestmentFilterGettingThousands valida el flujo cuando se debe obtener los miles de la cifra recibida
func TestInvestmentFilterGettingThousands(t *testing.T) {
	// Test data
	invest := int32(6700)
	credit1 := entities.CreditDetails{
		LoanQuantity: 300,
	}
	credit2 := entities.CreditDetails{
		LoanQuantity: 500,
	}
	credit3 := entities.CreditDetails{
		LoanQuantity: 700,
	}
	// Running Test
	service := NewInvestmentFilter()
	e := service.Filter(invest, &credit3, &credit1, &credit2)
	total := sumAllCredits(&credit1, &credit2, &credit3)

	// Test Validations
	assert.Nil(t, e, "No debió regresar error")
	assert.Equal(t, invest, total, "Las cantidades no coinciden!")
}

// TestInvestmentFilterNoAssignable valida el flujo cuando existe un remanente que no se puede asignar
func TestInvestmentFilterNoAssignable(t *testing.T) {
	// Test data
	invest := int32(400)
	credit1 := entities.CreditDetails{
		LoanQuantity: 300,
	}
	credit2 := entities.CreditDetails{
		LoanQuantity: 500,
	}
	credit3 := entities.CreditDetails{
		LoanQuantity: 700,
	}

	// Running Test
	service := NewInvestmentFilter()

	e := service.Filter(invest, &credit3, &credit1, &credit2)
	err, ok := e.(NoCreditAssigment)

	// Test Validations
	assert.NotNil(t, e, "No debió regresar error")
	assert.True(t, ok, "El error no es del tipo requerido")
	assert.Equal(t, invest, err.Investment, "El error no es del tipo requerido")

}

// TestIsLessThanZero valida el correcto funcionamiento de la función
func TestIsLessThanZero(t *testing.T) {

	dt := []struct {
		Data  int32
		Error bool
	}{
		{78, false},
		{-1, true},
	}

	for _, v := range dt {
		e := isLessThanZero(v.Data)
		if v.Error {
			assert.NotNil(t, e, fmt.Sprintf("Debió regresar error con el dato %d", v.Data))
		} else {
			assert.Nil(t, e, fmt.Sprintf("No debió regresar error con el dato %d", v.Data))
		}
	}

}
