package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var perso string = `[
	{  "id": 1,  "nombre": "Matias",  "apellido": "Perez",  "edad": 27 },
	{  "id": 2,  "nombre": "Juan",  "apellido": "Romero",  "edad": 35 }
	]`

type DummyLogger struct{}
type StubStore struct{}
type SpyStore struct {
	writeUsado bool
}

func (d *DummyLogger) Log(string) error {
	return nil
}

func (s *StubStore) Read(data interface{}) error {
	// Quiero que este Read haga lo mismo que el read del file.go pero sin realmente leer un fail (leer un archivo, base de datos, etc)
	file := []byte(perso)
	return json.Unmarshal(file, data)
}

func (s *StubStore) Write(data interface{}) error {
	// Como no la voy a usar, que solo devuelva nil, y ya cumple con la interface Store

	return nil
}

func (s *SpyStore) Read(data interface{}) error {
	// Quiero que este Read haga lo mismo que el read del file.go pero sin realmente leer un fail (leer un archivo, base de datos, etc)

	return nil
}

func (s *SpyStore) Write(data interface{}) error {
	// Como no la voy a usar, que solo devuelva nil, y ya cumple con la interface Store
	s.writeUsado = true
	return nil
}

func TestGetAll(t *testing.T) {
	db := StubStore{}
	r := NewRepository(&db)
	var dataExpected []Persona
	_ = json.Unmarshal([]byte(perso), &dataExpected)

	received, err := r.GetAll()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, dataExpected, received)
}

func TestUpdate(t *testing.T) {
	db := StubStore{}
	r := NewRepository(&db)
	//newT := `{  "id": 1,  "nombre": "Macarena",  "apellido": "Sosa",  "edad": 22 }`
	_, err := r.Update(1, "Macarena", "Sosa", 22)

	assert.Equal(t, err, nil)
}
func TestUpdateSpy(t *testing.T) {
	db := SpyStore{false}
	r := NewRepository(&db)
	//newT := `{  "id": 1,  "nombre": "Macarena",  "apellido": "Sosa",  "edad": 22 }`
	_, err := r.Update(1, "Macarena", "Sosa", 22)

	assert.Equal(t, err, nil)
	assert.True(t, db.writeUsado)
}

func TestSumarDummy(t *testing.T) {
	log := DummyLogger{}
	db := StubStore{}
	r := NewRepository(&db)
	expected := 8
	//newT := `{  "id": 1,  "nombre": "Macarena",  "apellido": "Sosa",  "edad": 22 }`
	suma := r.Sumar(3, 5, &log)

	assert.Equal(t, expected, suma)
}
