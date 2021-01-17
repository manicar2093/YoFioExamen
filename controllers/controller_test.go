package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/manicar2093/YoFioExamen/services"

	"github.com/manicar2093/YoFioExamen/models"

	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"
	"github.com/manicar2093/YoFioExamen/mocks"
)

// TestHandleCreditAssignment valida el flujo sin errores. Verifica que se envíe corretamente
// la información en formato application/json
func TestHandleCreditAssignment(t *testing.T) {

	testPath := "/credit-assignment"
	var askedAmount int32 = 3000
	x := fmt.Sprintf(`{"investment": %d}`, askedAmount)
	requestData := strings.NewReader(x)

	r := httptest.NewRequest(http.MethodPost, testPath, requestData)
	w := httptest.NewRecorder()

	investmentFilterMock := mocks.InvestmentFilter{}
	creditDetailsServiceMock := mocks.CreditDetailsServiceMock{}
	investmentFilterMock.On("Assign", askedAmount).Return(int32(2), int32(2), int32(2), nil)

	controller := NewCreditController(&investmentFilterMock, &creditDetailsServiceMock)

	s := mux.NewRouter()
	s.HandleFunc(testPath, controller.HandleCreditAssignment)
	s.ServeHTTP(w, r)

	var response models.CreditAssignResponse

	e := json.NewDecoder(w.Body).Decode(&response)
	if e != nil {
		t.Error("Error al obtener la respuesta", e)
	}

	investmentFilterMock.AssertExpectations(t)
	assert.Equal(t, 200, w.Code, "El código de respuesta debió ser 200")
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "No hay cabecero Content-Type requerido")
	assert.True(t, response.CreditType300 == 2, "No hay información dentro de la respuesta")
	assert.True(t, response.CreditType500 == 2, "No hay información dentro de la respuesta")
	assert.True(t, response.CreditType700 == 2, "No hay información dentro de la respuesta")

}

// TestHandleCreditAssignmentJsonDecodeError valida se mande un error 500 cuando el tipo de dato recibido
func TestHandleCreditAssignmentJsonDecodeError(t *testing.T) {

	testPath := "/credit-assignment"
	askedAmount := 3000
	x := fmt.Sprintf(`{"investment": "%d"}`, askedAmount)
	requestData := strings.NewReader(x)

	r := httptest.NewRequest(http.MethodPost, testPath, requestData)
	w := httptest.NewRecorder()

	investmentFilterMock := mocks.InvestmentFilter{}
	creditDetailsServiceMock := mocks.CreditDetailsServiceMock{}

	controller := NewCreditController(&investmentFilterMock, &creditDetailsServiceMock)

	s := mux.NewRouter()
	s.HandleFunc(testPath, controller.HandleCreditAssignment)
	s.ServeHTTP(w, r)

	investmentFilterMock.AssertExpectations(t)
	assert.Equal(t, 500, w.Code, "El código de respuesta debió ser 500")

}

// TestHandleCreditAssignmentWError valida el comportamiento cuando hay un error inesperado en
// el servicio
func TestHandleCreditAssignmentWError(t *testing.T) {

	testPath := "/credit-assignment"
	var askedAmount int32 = 3000
	x := fmt.Sprintf(`{"investment": %d}`, askedAmount)
	requestData := strings.NewReader(x)

	r := httptest.NewRequest(http.MethodPost, testPath, requestData)
	w := httptest.NewRecorder()

	investmentFilterMock := mocks.InvestmentFilter{}
	creditDetailsServiceMock := mocks.CreditDetailsServiceMock{}

	investmentFilterMock.On("Assign", askedAmount).Return(int32(0), int32(0), int32(0), errors.New("A random error"))

	controller := NewCreditController(&investmentFilterMock, &creditDetailsServiceMock)

	s := mux.NewRouter()
	s.HandleFunc(testPath, controller.HandleCreditAssignment)
	s.ServeHTTP(w, r)

	investmentFilterMock.AssertExpectations(t)
	assert.Equal(t, 500, w.Code, "El código de respuesta debió ser 500")
}

// TestHandleCreditAssignmentWNoCreditAssigment valida se envíe la respuesta necesaria cuando no
// se asignaron créditos con la inversión recibida
func TestHandleCreditAssignmentWNoCreditAssigment(t *testing.T) {

	testPath := "/credit-assignment"
	var askedAmount int32 = 400
	x := fmt.Sprintf(`{"investment": %d}`, askedAmount)
	requestData := strings.NewReader(x)

	r := httptest.NewRequest(http.MethodPost, testPath, requestData)
	w := httptest.NewRecorder()

	investmentFilterMock := mocks.InvestmentFilter{}
	creditDetailsServiceMock := mocks.CreditDetailsServiceMock{}

	investmentFilterMock.On("Assign", askedAmount).Return(int32(0), int32(0), int32(0), services.NoCreditAssigment{Investment: askedAmount, Remaining: askedAmount})

	controller := NewCreditController(&investmentFilterMock, &creditDetailsServiceMock)

	s := mux.NewRouter()
	s.HandleFunc(testPath, controller.HandleCreditAssignment)
	s.ServeHTTP(w, r)

	investmentFilterMock.AssertExpectations(t)
	assert.Equal(t, 400, w.Code, "El código de respuesta debió ser 500")
}

func TestHandleGetStatistics(t *testing.T) {

	testPath := "/statistics"

	r := httptest.NewRequest(http.MethodPost, testPath, nil)
	w := httptest.NewRecorder()

	creditAssignmentStatisticsMock := models.CreditsAssignmentStatistics{
		DoneAssignments:               10,
		SuccessfulAssignments:         90,
		UnsuccessfulAssignements:      88,
		AverageSuccessfulInvestment:   12,
		AverageUnsuccessfulInvestment: 11,
	}

	investmentFilterMock := mocks.InvestmentFilter{}
	creditDetailsServiceMock := mocks.CreditDetailsServiceMock{}

	creditDetailsServiceMock.On("GetStatistics").Return(
		creditAssignmentStatisticsMock, nil)

	controller := NewCreditController(&investmentFilterMock, &creditDetailsServiceMock)

	s := mux.NewRouter()
	s.HandleFunc(testPath, controller.HandleGetStatistics)
	s.ServeHTTP(w, r)

	var response models.CreditsAssignmentStatistics

	e := json.NewDecoder(w.Body).Decode(&response)
	if e != nil {
		t.Error("Error al obtener la respuesta", e)
	}

	investmentFilterMock.AssertExpectations(t)
	creditDetailsServiceMock.AssertExpectations(t)
	assert.Equal(t, 200, w.Code, "El código de respuesta debió ser 200")
	assert.Equal(t, creditAssignmentStatisticsMock.DoneAssignments, response.DoneAssignments, "El código de respuesta debió ser 200")
	assert.Equal(t, creditAssignmentStatisticsMock.SuccessfulAssignments, response.SuccessfulAssignments, "El código de respuesta debió ser 200")
	assert.Equal(t, creditAssignmentStatisticsMock.UnsuccessfulAssignements, response.UnsuccessfulAssignements, "El código de respuesta debió ser 200")
	assert.Equal(t, creditAssignmentStatisticsMock.AverageSuccessfulInvestment, response.AverageSuccessfulInvestment, "El código de respuesta debió ser 200")
	assert.Equal(t, creditAssignmentStatisticsMock.AverageUnsuccessfulInvestment, response.AverageUnsuccessfulInvestment, "El código de respuesta debió ser 200")
}

func TestHandleGetStatisticsGetStatisticsError(t *testing.T) {

	testPath := "/statistics"

	r := httptest.NewRequest(http.MethodPost, testPath, nil)
	w := httptest.NewRecorder()

	creditAssignmentStatisticsMock := models.CreditsAssignmentStatistics{}

	investmentFilterMock := mocks.InvestmentFilter{}
	creditDetailsServiceMock := mocks.CreditDetailsServiceMock{}

	creditDetailsServiceMock.On("GetStatistics").Return(
		creditAssignmentStatisticsMock, errors.New("An erros has occured"))

	controller := NewCreditController(&investmentFilterMock, &creditDetailsServiceMock)

	s := mux.NewRouter()
	s.HandleFunc(testPath, controller.HandleGetStatistics)
	s.ServeHTTP(w, r)

	investmentFilterMock.AssertExpectations(t)
	creditDetailsServiceMock.AssertExpectations(t)
	assert.Equal(t, 500, w.Code, "El código de respuesta debió ser 500")
}
