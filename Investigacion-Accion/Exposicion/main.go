package main

import (
	"errors"
	"fmt"
	"math"
)

var errorDivisionCero = errors.New("math: division por cero")
var errorRaizCuadrada = errors.New("math: raiz de un numero negativo")
var errorPunteroNulo = errors.New("pointer: puntero no referenciado")
var errorIndiceDesbordado = errors.New("index: acceso a un indice incorrecto")

func Dividir(num, den float64) (float64, error) {
	if den == 0.0 {
		return 0, errorDivisionCero
	}
	return num / den, nil
}
func RaizCuadrada(f float64) (float64, error) {
	if f < 0 {
		return 0, errorRaizCuadrada
	}
	return math.Sqrt(f), nil
}

type vector struct {
	x, y float64
}

func Suma(vec *vector) (float64, error) {
	if vec == nil {
		return 0, errorPunteroNulo
	}
	return vec.x + vec.y, nil
}

func Sumar(nums []float64, cant int) (float64, error) {
	suma := 0.0

	if len(nums) < cant {
		return 0.0, errorIndiceDesbordado
	}

	for i := 0; i < cant; i++ {
		suma += nums[i]
	}

	return suma, nil
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	// num1, num2 := 5.0, 10.0
	num1, num2 := 5.0, 10.0

	defer func() {
		res, err := Dividir(num1, num2)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Division", num1, "/", num2, "=", res)
		}
	}()

	defer func() {
		// num3 := -4.0
		num3 := -4.0
		res, err := RaizCuadrada(num3)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Raiz cuadrada", num3, "=", res)
		}
	}()

	defer func() {

		// var vec *vector
		vec := &vector{5, 6}

		res, err := Suma(vec)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Suma", vec.x, "+", vec.y, "=", res)
		}
	}()

	defer func() {
		arreglo := []float64{2.2, 3.4, 56}
		res, err := Sumar(arreglo, len(arreglo))
		// res, err = Sumar(arreglo, 4)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Suma slice", arreglo, "es:", res)
		}
	}()

}
