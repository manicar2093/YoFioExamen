package models

type CreditAssignResponse struct {
	CreditType300 int32 `json:"credit_type_300"`
	CreditType500 int32 `json:"credit_type_500"`
	CreditType700 int32 `json:"credit_type_700"`
}

type CreditAssignRequest struct {
	Investment int32 `json:"investment"`
}

type CreditsAssignmentStatistics struct {
	DoneAssignments               int32 `json:"done_assignments"`
	SuccessfulAssignments         int32 `json:"successful_assignments"`
	UnsuccessfulAssignements      int32 `json:"unsuccessful_assignments"`
	AverageSuccessfulInvestment   int32 `json:"average_successful_investment"`
	AverageUnsuccessfulInvestment int32 `json:"average_unsuccessful_investment"`
}