package calculadora

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Logger interface {
	Log(string) error
}

// Función que recibe dos enteros, un objeto del tipo logger y retorna la   suma resultante
func Sumar(num1, num2 int, logger Logger) int {
	err := logger.Log("Ingreso a Función Sumar")
	if err != nil {
		return -99999
	}
	return num1 + num2
}

// se crea un un struct dummyLogger
type dummyLogger struct{}

//  Se escriben las funciones necesarios para que dummyLogger cumpla con la interfaz que va a reemplazar (Logger)
func (d *dummyLogger) Log(string) error {
	return nil
}

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8
	// Se genera el objeto dummy a usar para satisfacer la necesidad de la función Sumar
	myDummy := &dummyLogger{}
	// Se ejecuta el test
	resultado := Sumar(num1, num2, myDummy)
	// Se validan los resultados aprovechando testify
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

// se crea un un struct stubLogger
type stubLogger struct{}

//  Se escribe las funciones necesarias para que stubLogger retorne exactamente lo que necesitamos
func (s *stubLogger) Log(string) error {
	return errors.New("error desde stub")
}

func TestSumarError(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := -99999
	// Se genera el objeto stub a usar para satisfacer la necesidad de la función Sumar
	myStub := &stubLogger{}
	// Se ejecuta el test
	resultado := Sumar(num1, num2, myStub)
	// Se validan los resultados aprovechando testify
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

// se crea un un struct spy compuesto por un booleano que nos informará si ocurre el llamado a Log
type spyLogger struct {
	spyCalled bool
}

//  Para espiar creamos un loggerSpy que setea en true spyCalled si entra al método
func (s *spyLogger) Log(string) error {
	s.spyCalled = true
	return nil
}
func TestSumarConSpy(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8
	// Se genera el objeto spy a usar
	mySpy := &spyLogger{}
	// Se ejecuta el test y se validan el resultado y que spyCalled sea true para dar el test por válido
	resultado := Sumar(num1, num2, mySpy)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
	assert.True(t, mySpy.spyCalled)
}

type Config interface {
	SumaEnabled(cliente string) bool
}

// Función que recibe dos enteros y retorna la suma resultante
func SumarRestricted(num1, num2 int, config Config, cliente string) int {
	if !config.SumaEnabled(cliente) {
		return -99999
	}
	return num1 + num2

}

// se crea un un struct mockConfig
type mockConfig struct {
	clienteUsado string
}

// El mock debe implementar el método necesario y comprobar que SumaEnabled sea llamado y que se haga exactamente con el mismo cliente que recibió SumarRestricted
func (m *mockConfig) SumaEnabled(cliente string) bool {
	m.clienteUsado = cliente
	return true
}
func TestSumarRestricted(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	cliente := "John Doe"
	resultadoEsperado := 8
	// Se genera el objeto dummy a usar para satisfacer la necesidad de la función Sumar
	myMock := &mockConfig{}
	// Se ejecuta el test y se valida el resultado y que el mock haya registrado la información correcta
	resultado := SumarRestricted(num1, num2, myMock, cliente)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
	assert.Equal(t, cliente, myMock.clienteUsado)
}

// Función que recibe dos enteros y retorna la suma resultante
func SumarRestrictedFake(num1, num2 int, config Config, cliente string) int {
	if !config.SumaEnabled(cliente) {
		return -99999
	}
	return num1 + num2
}

// se crea un un struct fakeConfig que implemente una lógica en la que sólo habilita la suma al cliente "John Doe"
type fakeConfig struct{}

func (f *fakeConfig) SumaEnabled(cliente string) bool {
	return cliente == "John Doe"
}
func TestSumarRestrictedFake(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	cliente := "John Doe"
	cliente_dos := "Mister Pmosh"
	resultadoEsperado := 8
	resultadoEsperadoError := -99999
	// Se genera el objeto fake a usar
	myFake := &fakeConfig{}
	// Se ejecuta el test y Se valida que para el cliente autorizado devuelva el resultado correcto de la suma y que para el cliente no autorizado devuelva el número -99999
	resultado := SumarRestrictedFake(num1, num2, myFake, cliente)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
	resultado2 := SumarRestrictedFake(num1, num2, myFake, cliente_dos)
	assert.Equal(t, resultadoEsperadoError, resultado2, "deben ser iguales")
}
