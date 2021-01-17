package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manicar2093/YoFioExamen/models"
	"github.com/manicar2093/YoFioExamen/utils"

	"github.com/manicar2093/YoFioExamen/services"
)

type CreditController interface {
	HandleCreditAssignment(w http.ResponseWriter, r *http.Request)
	HandleGetStatistics(w http.ResponseWriter, r *http.Request)
}

type CreditControllerImpl struct {
	creditAssigner       services.CreditAssigner
	creditDetailsService services.CreditDetailsService
}

func NewCreditController(creditAssigner services.CreditAssigner, creditDetailsService services.CreditDetailsService) CreditController {
	return &CreditControllerImpl{
		creditAssigner,
		creditDetailsService,
	}
}

func (c CreditControllerImpl) HandleCreditAssignment(w http.ResponseWriter, r *http.Request) {

	var body models.CreditAssignRequest
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
		utils.LogError.Println(fmt.Sprintf("El tipo de dato recibido no es correcto. \n\tDetalles: %v", e))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	loan1, loan2, loan3, e := c.creditAssigner.Assign(body.Investment)

	if e != nil {

		if _, ok := e.(services.NoCreditAssigment); ok {
			utils.LogError.Println(fmt.Sprintf("No se realizó asignación de creditos:\n\tDetalles: %v", e))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		utils.LogError.Println(fmt.Sprintf("Error inesperado al asignar creditos:\n\tDetalles: %v", e))
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	response := models.CreditAssignResponse{
		CreditType300: loan1,
		CreditType500: loan2,
		CreditType700: loan3,
	}

	utils.JSON(w, http.StatusOK, &response)
	return

}

func (c CreditControllerImpl) HandleGetStatistics(w http.ResponseWriter, r *http.Request) {

	data, e := c.creditDetailsService.GetStatistics()
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, data)

}
