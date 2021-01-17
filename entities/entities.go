package entities

import "fmt"

// CreditDetails es el holder donde se lleva la información del Credito que se requiere calcular
type CreditDetails struct {
	Count        int32 `bson:"count"`
	LoanQuantity int32 `bson:"loan_quantity"`
}

func (c CreditDetails) ToString() string {
	return fmt.Sprintf("CreditDetail{LoanQuantity:%d, Count: %d, TotalWithCount: %d}\n", c.LoanQuantity, c.Count, c.TotalWithCount())
}

func (c CreditDetails) TotalWithCount() int32 {
	return c.LoanQuantity * c.Count
}

type CreditDetailsWithStatus struct {
	Investment int32
	Credit1Count int32
	Credit2Count int32
	Credit3Count int32
	Status string
}

type CreditsAssignmentStatistics struct {
	DoneAssignments               int `bson:"done_assignments"`
	SuccessfulAssignments         int `bson:"successful_assignments"`
	UnsuccessfulAssignements      int `bson:"unsuccessful_assignments"`
	AverageSuccessfulInvestment   int `bson:"average_successful_investment"`
	AverageUnsuccessfulInvestment int `bson:"average_unsuccessful_investment"`
}
