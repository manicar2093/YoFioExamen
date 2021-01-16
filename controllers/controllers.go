package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/manicar2093/YoFioExamen/models"
	"github.com/manicar2093/YoFioExamen/utils"

	"github.com/manicar2093/YoFioExamen/services"
)

type CreditController struct {
	creditService services.CreditAssigner
}

func NewCreditController(creditService services.CreditAssigner) *CreditController {
	return &CreditController{
		creditService,
	}
}

func (c CreditController) HandleCreditAssignment(w http.ResponseWriter, r *http.Request) {

	var body models.CreditAssignRequest
	json.NewDecoder(r.Body).Decode(&body)
	loan1, loan2, loan3, e := c.creditService.Assign(body.Investment)

	if e != nil {

		if _, ok := e.(services.NoCreditAssigment); ok {

			utils.JSON(w, http.StatusBadRequest, nil)
			return
		}

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
