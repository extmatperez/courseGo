package ordenamientos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	nums := []int{2, 5, 6, 3}
	numsOrdenados := []int{2, 3, 5, 6}

	ordenar(nums)

	assert.Equal(t, numsOrdenados, nums, "No estan ordenados")

}

func TestDividir(t *testing.T) {
	num1, num2 := 5, 0

	_, err := Dividir(num1, num2)

	assert.NotNil(t, err)

}
