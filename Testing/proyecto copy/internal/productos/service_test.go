package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/w2GoPrueba/GoTesting/Ejemplo/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAllService(t *testing.T) {

	var expected []Producto
	json.Unmarshal([]byte(prod), &expected)
	dataJson := []byte(prod)
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

	dataJson := []byte(prod)
	dbMock := store.Mock{
		Data: dataJson,
	}

	testProducto := Producto{
		Nombre: "Cocina",
		Precio: 86500,
	}

	storeStub := store.FileStore{
		Mock: &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, _ := myService.Update(1, testProducto.Nombre, testProducto.Precio)
	assert.Equal(t, testProducto.Nombre, result.Nombre)
	assert.Equal(t, testProducto.Precio, result.Precio)
	assert.Equal(t, 1, result.ID)
}

func TestStore(t *testing.T) {

	dataJson := []byte(`[]`)
	dbMock := store.Mock{
		Data: dataJson,
	}

	testProducto := Producto{
		Nombre: "Cocina",
		Precio: 86500,
	}

	storeStub := store.FileStore{
		Mock: &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, _ := myService.Store(testProducto.Nombre, testProducto.Precio)
	assert.Equal(t, testProducto.Nombre, result.Nombre)
	assert.Equal(t, testProducto.Precio, result.Precio)
	assert.Equal(t, 1, result.ID)
}
