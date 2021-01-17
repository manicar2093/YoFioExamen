package services

import (
	"errors"
	"github.com/manicar2093/YoFioExamen/entities"
	"github.com/manicar2093/YoFioExamen/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"testing"
)

// TestGetAllCreditDetails valida se regresen los entities.CreditDetails necesarios
func TestGetAllCreditDetails(t *testing.T) {

	creditDetailsDaoMock := mocks.CreditDetailsDaoMock{}

	service := NewCreditDetailsService(&creditDetailsDaoMock)
	data, e := service.GetAllCreditDetails()

	assert.Nil(t, e, "No debió regresar error")
	assert.Equal(t, 3, len(data), "Los elementos requeridos deben ser 3")

}

// TestSaveSuccessfulRequest valida el funcionamiento sin inconvenientes del método
func TestSaveSuccessfulRequest(t *testing.T) {
	// Data Test
	invest := int32(20000)
	credit1 := entities.CreditDetails{}
	credit2 := entities.CreditDetails{}
	credit3 := entities.CreditDetails{}
	creditDetailsWStatus := entities.CreditDetailsWithStatus{
		Investment:     invest,
		CreditsDetails: []*entities.CreditDetails{&credit1, &credit2, &credit3},
		Status:         "successful",
	}

	// Data Mock
	creditDetailsDaoMock := mocks.CreditDetailsDaoMock{}
	creditDetailsDaoMock.On("Save", &creditDetailsWStatus).Return(nil)

	// Running Test
	service := NewCreditDetailsService(&creditDetailsDaoMock)
	e := service.SaveSuccessfulRequest(&credit1, &credit2, &credit3, invest)

	// Test Validations
	creditDetailsDaoMock.AssertExpectations(t)
	assert.Nil(t, e, "No debió regresar error.")
}

// TestSaveSuccessfulRequestSaveError valida el funcionamiento cuando el metodo de Save del dao falla
func TestSaveSuccessfulRequestSaveError(t *testing.T) {
	// Data Test
	invest := int32(20000)
	credit1 := entities.CreditDetails{}
	credit2 := entities.CreditDetails{}
	credit3 := entities.CreditDetails{}
	creditDetailsWStatus := entities.CreditDetailsWithStatus{
		Investment:     invest,
		CreditsDetails: []*entities.CreditDetails{&credit1, &credit2, &credit3},
		Status:         "successful",
	}

	// Data Mock
	creditDetailsDaoMock := mocks.CreditDetailsDaoMock{}
	creditDetailsDaoMock.On("Save", &creditDetailsWStatus).Return(errors.New("An error has occured"))

	// Running Test
	service := NewCreditDetailsService(&creditDetailsDaoMock)
	e := service.SaveSuccessfulRequest(&credit1, &credit2, &credit3, invest)

	// Test Validations
	creditDetailsDaoMock.AssertExpectations(t)
	assert.NotNil(t, e, "No debió regresar error.")
}

// TestSaveUnsuccessfulRequest valida el funcionamiento sin inconvenientes del método
func TestSaveUnsuccessfulRequest(t *testing.T) {
	// Data Test
	invest := int32(20000)
	credit1 := entities.CreditDetails{}
	credit2 := entities.CreditDetails{}
	credit3 := entities.CreditDetails{}
	creditDetailsWStatus := entities.CreditDetailsWithStatus{
		Investment:     invest,
		CreditsDetails: []*entities.CreditDetails{&credit1, &credit2, &credit3},
		Status:         "unsuccessful",
	}

	// Data Mock
	creditDetailsDaoMock := mocks.CreditDetailsDaoMock{}
	creditDetailsDaoMock.On("Save", &creditDetailsWStatus).Return(nil)

	// Running Test
	service := NewCreditDetailsService(&creditDetailsDaoMock)
	e := service.SaveUnsuccessfulRequest(&credit1, &credit2, &credit3, invest)

	// Test Validations
	creditDetailsDaoMock.AssertExpectations(t)
	assert.Nil(t, e, "No debió regresar error.")
}

// TestSaveUnsuccessfulRequestSaveError valida el funcionamiento cuando el metodo de Save del dao falla
func TestSaveUnsuccessfulRequestSaveError(t *testing.T) {
	// Data Test
	invest := int32(20000)
	credit1 := entities.CreditDetails{}
	credit2 := entities.CreditDetails{}
	credit3 := entities.CreditDetails{}
	creditDetailsWStatus := entities.CreditDetailsWithStatus{
		Investment:     invest,
		CreditsDetails: []*entities.CreditDetails{&credit1, &credit2, &credit3},
		Status:         "unsuccessful",
	}

	// Data Mock
	creditDetailsDaoMock := mocks.CreditDetailsDaoMock{}
	creditDetailsDaoMock.On("Save", &creditDetailsWStatus).Return(errors.New("An error has occured"))

	// Running Test
	service := NewCreditDetailsService(&creditDetailsDaoMock)
	e := service.SaveUnsuccessfulRequest(&credit1, &credit2, &credit3, invest)

	// Test Validations
	creditDetailsDaoMock.AssertExpectations(t)
	assert.NotNil(t, e, "No debió regresar error.")
}

func TestGetStatistics(t *testing.T) {

	creditsSuccess := []entities.CreditDetailsWithStatus{
		{
			Investment:     30000,
			CreditsDetails: nil,
			Status:         "successful",
		},
	}
	creditsUnsuccessful := []entities.CreditDetailsWithStatus{
		{
			Investment:     15000,
			CreditsDetails: nil,
			Status:         "unsuccessful",
		},
	}

	creditDetailsDaoMock := mocks.CreditDetailsDaoMock{}
	creditDetailsDaoMock.On("FilterCreditDetailsWithStatus", bson.D{
		primitive.E{
			Key:   "status",
			Value: "successful",
		},
	}).Return(creditsSuccess, nil)

	creditDetailsDaoMock.On("FilterCreditDetailsWithStatus", bson.D{
		primitive.E{
			Key:   "status",
			Value: "unsuccessful",
		},
	}).Return(creditsUnsuccessful, nil)

	service := NewCreditDetailsService(&creditDetailsDaoMock)
	data, e := service.GetStatistics()

	creditDetailsDaoMock.AssertExpectations(t)
	assert.Nil(t, e, "No debió regresar error")
	assert.Equal(t, int32(2), data.DoneAssignments, "No debe traer contadore DoneAssignments")
	assert.Equal(t, int32(1), data.SuccessfulAssignments, "No debe traer contadore SuccessfulAssignments")
	assert.Equal(t, int32(1), data.UnsuccessfulAssignements, "No debe traer contadore UnsuccessfulAssignements")
	assert.Equal(t, int32(30000), data.AverageSuccessfulInvestment, "No debe traer contadore AverageSuccessfulInvestment")
	assert.Equal(t, int32(15000), data.AverageUnsuccessfulInvestment, "No debe traer contadore AverageUnsuccessfulInvestment")
}
