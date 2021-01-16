package dao

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllCreditDetails(t *testing.T) {
	dao := NewCreditDetailsDaoImpl()
	data, e := dao.GetAllCreditDetails()
	assert.Nil(t, e, "No debió regresar error")
	assert.Equal(t, 3, len(data), "No se recibió la cantidad de datos requeridos")
}