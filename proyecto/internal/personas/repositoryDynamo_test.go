package internal

import (
	"context"
	"testing"

	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_dynamoRepository_Store(t *testing.T) {
	dynamo, err := InitDynamo()
	assert.NoError(t, err)
	repository := NewDynamoRepository(dynamo, "Personas")
	ctx := context.TODO()
	userId := "9"
	user := models.PersonaDynamo{
		ID:       userId,
		Nombre:   "Matias",
		Apellido: "Perez",
	}
	err = repository.Store(ctx, &user)
	assert.NoError(t, err)
	getResult, err := repository.GetOne(ctx, userId)
	assert.NoError(t, err)
	assert.NotNil(t, getResult)
	assert.Equal(t, user.ID, getResult.ID)
	//	err = repository.Delete(ctx, userId)
	//	assert.NoError(t, err)
}
