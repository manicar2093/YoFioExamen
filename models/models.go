package models

import "fmt"

type CreditAssignResponse struct {
	CreditType300 int32 `json:"credit_type_300"`
	CreditType500 int32 `json:"credit_type_500"`
	CreditType700 int32 `json:"credit_type_700"`
}

type CreditAssignRequest struct {
	Investment int32 `json:"investment"`
}

// NoCreditAssigment indica que la inversión no se puede asignar a los creditos solicitados
type NoCreditAssigment struct {
	Investment int32
	Remaining  int32
}

func (n NoCreditAssigment) Error() string {
	return fmt.Sprintf("Se tiene un remanente de $%d al asignar la inversión de $%d", n.Remaining, n.Investment)
}