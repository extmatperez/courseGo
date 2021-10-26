package ordenamientos

import (
	"crypto/sha1"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	nums := []int{2, 5, 6, 3}
	numsOrdenados := []int{2, 3, 5, 6}

	ordenar(nums)

	assert.Equal(t, numsOrdenados, nums, "No estan ordenados")

}

func TestDividirError(t *testing.T) {
	num1, num2 := 5, 0

	_, err := Dividir(num1, num2)

	assert.NotNil(t, err)

}

func TestDividir(t *testing.T) {
	num1, num2 := 6, 2

	num, _ := Dividir(num1, num2)

	assert.Equal(t, 3, num)

}

func TestSumar(t *testing.T) {
	num1, num2 := 5, 0

	num := Sumar(num1, num2)

	assert.Equal(t, 5, num)

}
func TestRestar(t *testing.T) {
	num1, num2 := 5, 0

	num := Restar(num1, num2)

	assert.Equal(t, 5, num)

}

func BenchmarkSum256(b *testing.B) {
	data := []byte("Digital House impulsando la transformacion digital")
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}
func BenchmarkSum(b *testing.B) {
	data := []byte("Digital House impulsando la transformacion digital")
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}
