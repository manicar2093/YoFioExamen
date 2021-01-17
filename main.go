package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/manicar2093/YoFioExamen/controllers"
	"github.com/manicar2093/YoFioExamen/dao"
	"github.com/manicar2093/YoFioExamen/services"
	"github.com/manicar2093/YoFioExamen/utils"
	"log"
	"net/http"
)

var creditAssignerController controllers.CreditController
var creditAssigner services.CreditAssigner
var creditFilter services.InvestmentFilter
var crediteDetailsDao dao.CreditDetailsDao

func main() {
	fmt.Println(utils.GetBanner())

	r := mux.NewRouter()
	r.HandleFunc("/credit-assignment", creditAssignerController.HandleCreditAssignment).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8000", r))

}

func init() {
	//crediteDetailsDao = dao.NewCreditDetailsDaoImpl()
	//creditFilter = services.NewInvestmentFilter()
	//creditAssigner = services.NewCreditAssigner(creditFilter, crediteDetailsDao)
	//creditAssignerController = controllers.NewCreditController(creditAssigner)
}
