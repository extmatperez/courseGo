package calculadora

// Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8

	// Se ejecuta el test
	resultado := Sumar(num1, num2)

	// Se validan los resultados
	if resultado != resultadoEsperado {
		t.Errorf("Funcion suma() arrojo el resultado = %v, pero el esperado es  %v", resultado, resultadoEsperado)
	}

}

//go get github.com/stretchr/testify

func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 13
	num2 := 5
	resultadoEsperado := 8

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}

func TestVarios(t *testing.T) {
	assert.Equal(t, 123, 123, "deberían ser iguales")
	assert.NotEqual(t, 123, 456, "no deberían ser iguales")
	assert.Nil(t, nil, "Esperaba nulo")

	// assert.NotNil(t, nil, "No deberia ser nulo")
}

func TestDividir(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 2

	// Se ejecuta el test
	resultado := Dividir(num1, num2)

	// Se validan los resultados aprovechando testify
	assert.NotNil(t, resultado)

}
