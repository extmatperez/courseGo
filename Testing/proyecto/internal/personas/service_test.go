package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/w2GoPrueba/GoTesting/Clase2TT/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAllService(t *testing.T) {

	var expected []Persona
	json.Unmarshal([]byte(perso), &expected)
	dataJson := []byte(perso)
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestServiceGetAllError(t *testing.T) {
	expectedError := errors.New("error for GetAll")
	dbMock := store.Mock{
		Err: expectedError,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
}

func TestUpdate(t *testing.T) {

	dataJson := []byte(perso)
	dbMock := store.Mock{
		Data: dataJson,
	}

	testPersona := Persona{
		Nombre:   "Cristian",
		Apellido: "Lopez",
		Edad:     36,
	}

	storeStub := store.FileStore{
		Mock: &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, _ := myService.Update(1, testPersona.Nombre, testPersona.Apellido, testPersona.Edad)
	assert.Equal(t, testPersona.Nombre, result.Nombre)
	assert.Equal(t, testPersona.Apellido, result.Apellido)
	assert.Equal(t, testPersona.Edad, result.Edad)
	assert.Equal(t, 1, result.ID)
}

func TestStore(t *testing.T) {

	dataJson := []byte(`[]`)
	dbMock := store.Mock{
		Data: dataJson,
	}

	testPersona := Persona{
		Nombre:   "Cristian",
		Apellido: "Lopez",
		Edad:     36,
	}

	storeStub := store.FileStore{
		Mock: &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, _ := myService.Store(testPersona.Nombre, testPersona.Apellido, testPersona.Edad)
	assert.Equal(t, testPersona.Nombre, result.Nombre)
	assert.Equal(t, testPersona.Apellido, result.Apellido)
	assert.Equal(t, testPersona.Edad, result.Edad)
	assert.Equal(t, 1, result.ID)
}
