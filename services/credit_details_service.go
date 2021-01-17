package services

import (
	"github.com/manicar2093/YoFioExamen/dao"
	"github.com/manicar2093/YoFioExamen/entities"
	"github.com/manicar2093/YoFioExamen/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	Successful   = "successful"
	Unsuccessful = "unsuccessful"
)

type CreditDetailsService interface {
	// GetAllCreditDetails obtiene todos los entities.CreditDetails que se deben validar. Existe considerando que puede extraerse esta información de una base de datos.
	// Solo se debe implementar la lógica necesaria
	GetAllCreditDetails() ([]entities.CreditDetails, error)
	SaveSuccessfulRequest(credit1 *entities.CreditDetails, credit2 *entities.CreditDetails, credit3 *entities.CreditDetails, invest int32) error
	SaveUnsuccessfulRequest(credit1 *entities.CreditDetails, credit2 *entities.CreditDetails, credit3 *entities.CreditDetails, invest int32) error
	// GetStatistics realiza la busquda de los CreditDetails que se han registrado y realiza el calculo estadistico
	GetStatistics() (models.CreditsAssignmentStatistics, error)
}

type CreditDetailsServiceImpl struct {
	creditDetailsDao dao.CreditDetailsDao
}

func NewCreditDetailsService(creditDetailsDao dao.CreditDetailsDao) CreditDetailsService {
	return &CreditDetailsServiceImpl{creditDetailsDao: creditDetailsDao}
}

func (c CreditDetailsServiceImpl) GetAllCreditDetails() ([]entities.CreditDetails, error) {
	return []entities.CreditDetails{
		{LoanQuantity: 300, Count: 0},
		{LoanQuantity: 500, Count: 0},
		{LoanQuantity: 700, Count: 0},
	}, nil
}

func (c CreditDetailsServiceImpl) SaveSuccessfulRequest(credit1 *entities.CreditDetails, credit2 *entities.CreditDetails, credit3 *entities.CreditDetails, invest int32) error {
	return c.creditDetailsDao.Save(createCreditDetailsWStatus(credit1, credit2, credit3, invest, Successful))
}

func (c CreditDetailsServiceImpl) SaveUnsuccessfulRequest(credit1 *entities.CreditDetails, credit2 *entities.CreditDetails, credit3 *entities.CreditDetails, invest int32) error {
	return c.creditDetailsDao.Save(createCreditDetailsWStatus(credit1, credit2, credit3, invest, Unsuccessful))
}

func (c CreditDetailsServiceImpl) GetStatistics() (models.CreditsAssignmentStatistics, error) {
	calculo := models.CreditsAssignmentStatistics{}
	success, _ := c.creditDetailsDao.FilterCreditDetailsWithStatus(primitive.D{
		primitive.E{
			Key:   "status",
			Value: Successful,
		},
	})
	unsuccess, _ := c.creditDetailsDao.FilterCreditDetailsWithStatus(primitive.D{
		primitive.E{
			Key:   "status",
			Value: Unsuccessful,
		},
	})

	calculateAverage(&success, &calculo.AverageSuccessfulInvestment, &calculo.SuccessfulAssignments)
	calculateAverage(&unsuccess, &calculo.AverageUnsuccessfulInvestment, &calculo.UnsuccessfulAssignements)

	calculo.DoneAssignments = calculo.UnsuccessfulAssignements + calculo.SuccessfulAssignments
	return calculo, nil

}

func createCreditDetailsWStatus(credit1 *entities.CreditDetails, credit2 *entities.CreditDetails, credit3 *entities.CreditDetails, invest int32, status string) *entities.CreditDetailsWithStatus {
	return &entities.CreditDetailsWithStatus{
		Investment:     invest,
		CreditsDetails: []*entities.CreditDetails{credit1, credit2, credit3},
		Status:         status,
	}
}

// calculateAverage realiza el calculo del promedio. El averangePointer es donde se debe asignar el dato del promedio; un campo Average***Investment.
// El assignmentCount es donde se debe asignar el dato de la cantidad; un campo ****Assignments
func calculateAverage(creditDetails *[]entities.CreditDetailsWithStatus, averagePointer, assignmentCount *int32) {
	var totalAmount, detailsCount int32 = 0, int32(len(*creditDetails))

	for _, v := range *creditDetails {
		totalAmount += v.Investment
	}

	if totalAmount == 0 {
		*averagePointer = 0
		*assignmentCount = detailsCount
		return
	}

	*averagePointer = totalAmount / detailsCount
	*assignmentCount = detailsCount

}
