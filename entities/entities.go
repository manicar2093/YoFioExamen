package entities

import "fmt"

// CreditDetails es el holder donde se lleva la información del Credito que se requiere calcular
type CreditDetails struct {
	Count        int32
	LoanQuantity int32
}

func (c CreditDetails) ToString() string {
	return fmt.Sprintf("CreditDetail{LoanQuantity:%d, Count: %d, TotalWithCount: %d}\n", c.LoanQuantity, c.Count, c.TotalWithCount())
}

func (c CreditDetails) TotalWithCount() int32 {
	return c.LoanQuantity * c.Count
}