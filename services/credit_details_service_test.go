package services

import (

	"github.com/manicar2093/YoFioExamen/mocks"
	"github.com/stretchr/testify/assert"

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
/*
func TestGetStatistics(t *testing.T) {

	creditDetailsDaoMock := mocks.CreditDetailsDaoMock{}
	creditDetailsDaoMock.On("FindStatistics").Return(entities.CreditsAssignmentStatistics{}, nil)
	creditDetailsDaoMock.On("FindCreditCounters", bson.D{primitive.E{
			Key:   "status",
			Value: "successful",
		}}).Return([]entities.CreditDetailsWithStatus{{},{},{},}, nil)

	creditDetailsDaoMock.On("FindCreditCounters", bson.D{primitive.E{
			Key:   "status",
			Value: "unsuccessful",
		}}).Return([]entities.CreditDetailsWithStatus{{},{},{},}, nil)

	service := NewCreditDetailsService(&creditDetailsDaoMock)
	data, e := service.GetStatistics()

	creditDetailsDaoMock.AssertExpectations(t)
	assert.Nil(t, e, "No debió regresar error")
	assert.Equal(t, 0, data.DoneAssignments, "No debe traer contadore DoneAssignments")
	assert.Equal(t, 0, data.SuccessfulAssignments, "No debe traer contadore SuccessfulAssignments")
	assert.Equal(t, 0, data.UnsuccessfulAssignements, "No debe traer contadore UnsuccessfulAssignements")
	assert.Equal(t, 0, data.AverageSuccessfulInvestment, "No debe traer contadore AverageSuccessfulInvestment")
	assert.Equal(t, 0, data.AverageUnsuccessfulInvestment, "No debe traer contadore AverageUnsuccessfulInvestment")
}
*/