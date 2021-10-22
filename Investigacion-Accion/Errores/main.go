package main

import (
	"errors"
	"fmt"
	"math"
)

func RaizCuadrada(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: Numero negativo")
	}
	return math.Sqrt(f), nil
}

func saludar(mensaje string) {
	fmt.Println(mensaje)
}

func DoblePositivo(num int) (int, error, error) {
	if num == 0 {
		return 0, errors.New("el numero no puede ser 0"), nil
	} else {
		if num < 0 {
			return 0, nil, errors.New("el numero no puede ser negativo")
		}
	}
	return num * 2, nil, nil
}

var errorNegativo = errors.New("el numero no puede ser negativo")
var errorCero = errors.New("el numero no puede ser 0")

type Wrapper interface {
	Unwrap() error
}

func DoblePositivoVersionDos(num int) (int, error) {
	if num == 0 {
		return 0, errorCero
	} else {
		if num < 0 {
			return 0, errorNegativo
		}
	}
	return num * 2, nil
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic recuperado:", err)
		}
	}()

	num := 5
	doble, err := DoblePositivoVersionDos(num)

	switch err {
	case nil:
		fmt.Printf("El doble de %d es %d", num, doble)
	case errorCero:
		fmt.Println("Case cero")
		panic(err)
	case errorNegativo:
		fmt.Println("Case negativo")
		panic(err)
	}

	if err != nil {
		panic(err)
	} else {
		fmt.Println(num)
	}

}

/*
defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	defer saludar("Empezando")

	num, err := RaizCuadrada(5)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(num)
	}

	defer saludar("Terminando")

*/
