package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubStore struct{}

func (s *stubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(`[{"id": 1,"nombre": "Ibby","apellido": "Rabat","email": "irabat0@digg.com","edad": 46,"altura": 1.49,"activo": false,"fechaDeCreacion": "1/15/2021"},
		{"id": 2,"nombre": "Ibby","apellido": "Rabat","email": "irabat0@digg.com","edad": 46,"altura": 1.49,"activo": false,"fechaDeCreacion": "1/15/2021"}]`),
		&data)
}

func (s *stubStore) Write(data interface{}) error {
	return nil
}

func TestStub(t *testing.T) {
	var expectedUser []Usuario
	json.Unmarshal([]byte(`[{"id": 1,"nombre": "Ibby","apellido": "Rabat","email": "irabat0@digg.com","edad": 46,"altura": 1.49,"activo": false,"fechaDeCreacion": "1/15/2021"},
		{"id": 2,"nombre": "Ibby","apellido": "Rabat","email": "irabat0@digg.com","edad": 46,"altura": 1.49,"activo": false,"fechaDeCreacion": "1/15/2021"}]`),
		&expectedUser)

	stub := &stubStore{}

	repo := NewRepository(stub)

	// service := NewService(repo)

	// salida, err := service.GetAll()
	salida, err := repo.GetAll()

	assert.Equal(t, expectedUser, salida, "Deberian ser iguales")
	assert.Nil(t, err, "Existe un error")

}

type mockStore struct {
	used bool
}

func (ms *mockStore) Read(data interface{}) error {
	ms.used = true
	return json.Unmarshal([]byte(`[{"id": 1,"nombre": "Before Update","apellido": "Rabat","email": "irabat0@digg.com","edad": 46,"altura": 1.49,"activo": false,"fechaDeCreacion": "1/15/2021"}]`), &data)
}

func (ms *mockStore) Write(data interface{}) error {
	return nil
}

func TestUpdateName(t *testing.T) {
	var expectedUser Usuario
	json.Unmarshal([]byte(`{"id": 1,"nombre": "After Update","apellido": "Rabat","email": "irabat0@digg.com","edad": 46,"altura": 1.49,"activo": false,"fechaDeCreacion": "1/15/2021"}`), &expectedUser)

	mock := &mockStore{used: false}

	repo := NewRepository(mock)

	prod, err := repo.UpdateNombre(1, "After Update")

	assert.Equal(t, expectedUser, prod, "No se actualizo correctamente")

	assert.True(t, mock.used, "No se actualizo")

	assert.Nil(t, err, "Hubo un error")
}
