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
}

type CreditControllerImpl struct {
	creditService services.CreditAssigner
}

func NewCreditController(creditService services.CreditAssigner) CreditController {
	return &CreditControllerImpl{
		creditService,
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
	loan1, loan2, loan3, e := c.creditService.Assign(body.Investment)

	if e != nil {

		if _, ok := e.(services.NoCreditAssigment); ok {
			utils.LogError.Println(fmt.Sprintf("No se realizó asignación de creditos:\n\tDetalles: %v", e))
			utils.JSON(w, http.StatusBadRequest, nil)
			return
		}
		utils.LogError.Println(fmt.Sprintf("Error inesperado al asignar creditos:\n\tDetalles: %v", e))
		utils.JSON(w, http.StatusInternalServerError, nil)
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
