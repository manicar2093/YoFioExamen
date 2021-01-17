package services

import (
	"errors"
	"github.com/manicar2093/YoFioExamen/entities"
	"github.com/manicar2093/YoFioExamen/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestAssign valida el flujo cuando no se presenta ningun inconveniente
func TestAssign(t *testing.T) {
	// Test data
	invest := int32(17800)
	credit300 := entities.CreditDetails{LoanQuantity: 300, Count: 5}
	credit500 := entities.CreditDetails{LoanQuantity: 500, Count: 5}
	credit700 := entities.CreditDetails{LoanQuantity: 700, Count: 5}

	// Test Mocks
	creditDetailsServiceMock := mocks.CreditDetailsServiceMock{}
	creditDetailsServiceMock.On("GetAllCreditDetails").
		Return([]entities.CreditDetails{
			credit300,
			credit500,
			credit700,
		}, nil)
	creditDetailsServiceMock.On("SaveSuccessfulRequest",
		&credit300,
		&credit500,
		&credit700, invest).Return(nil)

	filterMock := mocks.InvestmentFilterMock{}
	filterMock.On("Filter", invest,
		&credit300,
		&credit500,
		&credit700).Return(nil)

	// Running Test
	service := NewCreditAssigner(&filterMock, &creditDetailsServiceMock)
	c300, c500, c700, e := service.Assign(invest)

	// Test Validations
	filterMock.AssertExpectations(t)
	creditDetailsServiceMock.AssertExpectations(t)
	assert.Nil(t, e, "No debió regresar error")
	assert.Equal(t, credit300.Count, c300, "Mal conteo de los creditos de 300")
	assert.Equal(t, credit500.Count, c500, "Mal conteo de los creditos de 500")
	assert.Equal(t, credit700.Count, c700, "Mal conteo de los creditos de 700")

}

// TestAssignCreditDetailsDaoError valida el manejo de error cuando hay problemas al extraer los entities.CreditDetails
func TestAssignCreditDetailsDaoError(t *testing.T) {
	// Test data
	invest := int32(17800)

	// Test Mocks
	creditDetailsServiceMock := mocks.CreditDetailsServiceMock{}
	creditDetailsServiceMock.On("GetAllCreditDetails").
		Return([]entities.CreditDetails{}, errors.New("An error has occured"))

	filterMock := mocks.InvestmentFilterMock{}

	// Running Test
	service := NewCreditAssigner(&filterMock, &creditDetailsServiceMock)
	c300, c500, c700, e := service.Assign(invest)

	// Test Validations
	filterMock.AssertExpectations(t)
	creditDetailsServiceMock.AssertExpectations(t)
	assert.NotNil(t, e, "Debió regresar error")
	assert.Equal(t, int32(0), c300, "Mal conteo de los creditos de 300")
	assert.Equal(t, int32(0), c500, "Mal conteo de los creditos de 500")
	assert.Equal(t, int32(0), c700, "Mal conteo de los creditos de 700")

}

// TestAssignFilterError valida el manejo de error cuando hay un error en Filter
func TestAssignFilterError(t *testing.T) {
	// Test data
	invest := int32(17800)
	credit300 := entities.CreditDetails{LoanQuantity: 300, Count: 5}
	credit500 := entities.CreditDetails{LoanQuantity: 500, Count: 5}
	credit700 := entities.CreditDetails{LoanQuantity: 700, Count: 5}

	// Test Mocks
	creditDetailsServiceMock := mocks.CreditDetailsServiceMock{}
	creditDetailsServiceMock.On("GetAllCreditDetails").
		Return([]entities.CreditDetails{
			credit300,
			credit500,
			credit700,
		}, nil)

	creditDetailsServiceMock.On("SaveUnsuccessfulRequest",
		&credit300,
		&credit500,
		&credit700, invest).Return(nil)

	filterMock := mocks.InvestmentFilterMock{}
	filterMock.On("Filter", invest,
		&credit300,
		&credit500,
		&credit700).Return(NoCreditAssigment{Investment: invest, Remaining: 200})

	// Running Test
	service := NewCreditAssigner(&filterMock, &creditDetailsServiceMock)
	_, _, _, e := service.Assign(invest)
	_, ok := e.(NoCreditAssigment)

	// Test Validations
	filterMock.AssertExpectations(t)
	creditDetailsServiceMock.AssertExpectations(t)

	assert.NotNil(t, e, "No debió regresar error")
	assert.True(t, ok, "El error no es del tipo necesario")

}

// TestAssignSaveSuccessfulRequestError valida el flujo cuando falla el guardado de la información del request
func TestAssignSaveSuccessfulRequestError(t *testing.T) {
	// Test data
	invest := int32(17800)
	credit300 := entities.CreditDetails{LoanQuantity: 300, Count: 5}
	credit500 := entities.CreditDetails{LoanQuantity: 500, Count: 5}
	credit700 := entities.CreditDetails{LoanQuantity: 700, Count: 5}

	// Test Mocks
	creditDetailsServiceMock := mocks.CreditDetailsServiceMock{}
	creditDetailsServiceMock.On("GetAllCreditDetails").
		Return([]entities.CreditDetails{
			credit300,
			credit500,
			credit700,
		}, nil)
	creditDetailsServiceMock.On("SaveSuccessfulRequest",
		&credit300,
		&credit500,
		&credit700, invest).Return(errors.New("An error has occured"))

	filterMock := mocks.InvestmentFilterMock{}
	filterMock.On("Filter", invest,
		&credit300,
		&credit500,
		&credit700).Return(nil)

	// Running Test
	service := NewCreditAssigner(&filterMock, &creditDetailsServiceMock)
	c300, c500, c700, e := service.Assign(invest)

	// Test Validations
	filterMock.AssertExpectations(t)
	creditDetailsServiceMock.AssertExpectations(t)
	assert.NotNil(t, e, "No debió regresar error")
	assert.Equal(t, int32(0), c300, "Mal conteo de los creditos de 300")
	assert.Equal(t, int32(0), c500, "Mal conteo de los creditos de 500")
	assert.Equal(t, int32(0), c700, "Mal conteo de los creditos de 700")

}

// TestAssignSaveUnsuccessfulRequestError valida el flujo cuando falla el guardado de la información del request
func TestAssignSaveUnsuccessfulRequestError(t *testing.T) {
	// Test data
	invest := int32(17800)
	credit300 := entities.CreditDetails{LoanQuantity: 300, Count: 5}
	credit500 := entities.CreditDetails{LoanQuantity: 500, Count: 5}
	credit700 := entities.CreditDetails{LoanQuantity: 700, Count: 5}

	// Test Mocks
	creditDetailsServiceMock := mocks.CreditDetailsServiceMock{}
	creditDetailsServiceMock.On("GetAllCreditDetails").
		Return([]entities.CreditDetails{
			credit300,
			credit500,
			credit700,
		}, nil)

	creditDetailsServiceMock.On("SaveUnsuccessfulRequest",
		&credit300,
		&credit500,
		&credit700, invest).Return(errors.New("An error has ocurred"))

	filterMock := mocks.InvestmentFilterMock{}
	filterMock.On("Filter", invest,
		&credit300,
		&credit500,
		&credit700).Return(NoCreditAssigment{Investment: invest, Remaining: 200})

	// Running Test
	service := NewCreditAssigner(&filterMock, &creditDetailsServiceMock)
	_, _, _, e := service.Assign(invest)

	// Test Validations
	filterMock.AssertExpectations(t)
	creditDetailsServiceMock.AssertExpectations(t)

	assert.NotNil(t, e, "No debió regresar error")

}