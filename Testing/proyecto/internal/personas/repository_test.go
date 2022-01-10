package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/extmatperez/w2GoPrueba/GoTesting/Clase2TT/proyecto/pkg/store"
)

type StubStore struct{}

var perso string = `[ 
	{	"id": 1,	"nombre": "Matias",	"apellido": "Perez",	"edad": 27   },
   	{	"id": 2,	"nombre": "Juan",	"apellido": "Romero",	"edad": 25   }]`

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(perso), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAllMock(t *testing.T) {
	// Initializing input/output

	dataJson := []byte(perso)
	dbMock := store.Mock{
		Data: dataJson,
	}

	var expected []Persona
	json.Unmarshal([]byte(perso), &expected)
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
