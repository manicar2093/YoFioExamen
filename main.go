package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/manicar2093/YoFioExamen/connections"
	"github.com/manicar2093/YoFioExamen/controllers"
	"github.com/manicar2093/YoFioExamen/dao"
	"github.com/manicar2093/YoFioExamen/services"
	"github.com/manicar2093/YoFioExamen/utils"
	"log"
	"net/http"
	"time"
)

var creditAssignerController controllers.CreditController
var creditDetailService services.CreditDetailsService
var creditAssigner services.CreditAssigner
var creditFilter services.InvestmentFilter
var creditDetailsDao dao.CreditDetailsDao

func main() {
	fmt.Println(utils.GetBanner())
	r := mux.NewRouter()
	r.HandleFunc("/credit-assignment", creditAssignerController.HandleCreditAssignment).Methods(http.MethodPost)
	r.HandleFunc("/statistics", creditAssignerController.HandleGetStatistics).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8000", r))

}

func init() {
	db := connections.GetMongoConnection(context.WithTimeout(context.TODO(), 5*time.Second))
	creditDetailsDao = dao.NewCreditDetailsDao(db.Collection("credit_details"))

	creditDetailService = services.NewCreditDetailsService(creditDetailsDao)
	creditFilter = services.NewInvestmentFilter()
	creditAssigner = services.NewCreditAssigner(creditFilter, creditDetailService)

	creditAssignerController = controllers.NewCreditController(creditAssigner, creditDetailService)
}
