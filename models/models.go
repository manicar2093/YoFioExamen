package models

type CreditAssignResponse struct {
	CreditType300 int32 `json:"credit_type_300"`
	CreditType500 int32 `json:"credit_type_500"`
	CreditType700 int32 `json:"credit_type_700"`
}

type CreditAssignRequest struct {
	Investment int32 `json:"investment"`
}