package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/extmatperez/w2GoPrueba/GoTesting/Ejemplo/proyecto/pkg/store"
)

type StubStore struct{}

var prod string = `[ 
	{	"id": 1,	"nombre": "Heladera",	"precio": 65000},
   	{	"id": 2,	"nombre": "Mesa",	"precio":38000  }]`

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(prod), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAllMock(t *testing.T) {
	// Initializing input/output

	dataJson := []byte(prod)
	dbMock := store.Mock{
		Data: dataJson,
	}

	var expected []Producto
	json.Unmarshal([]byte(prod), &expected)
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	// Test Execution
	resp, _ := myRepo.GetAll()
	// Validation
	assert.Equal(t, expected, resp)
}
