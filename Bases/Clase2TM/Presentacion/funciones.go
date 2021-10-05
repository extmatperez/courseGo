package main

import (
	"errors"
	"fmt"
)

const (
	suma   = "+"
	resta  = "-"
	multip = "*"
	divis  = "/"
)

func sumarNumeros(num1, num2 int) int {
	return num1 + num2
}
func multiplicarNumeros(num1 int, num2 int) int {
	return num1 * num2
}

func operacionAritmetica(valor1, valor2 float64, operador string) float64 {
	switch operador {
	case suma:
		return valor1 + valor2
	case resta:
		return valor1 - valor2
	case multip:
		return valor1 * valor2
	case divis:
		if valor2 != 0 {
			return valor1 / valor2
		}
	}
	return 0
}

func sumaEllipsis(values ...int) int {
	fmt.Println("\nNumeros que llegan: ", values)

	var resultado int
	for _, value := range values {
		resultado += value
	}
	return resultado
}

func operaciones(valor1, valor2 float64) (float64, float64, float64, float64) {
	suma := valor1 + valor2
	resta := valor1 - valor2
	multip := valor1 * valor2
	var divis float64

	if valor2 != 0 {
		divis = valor1 / valor2
	}

	return suma, resta, multip, divis
}

func main() {

	a, b, c, d := 5, -5, 0, 25

	validarNumero(a)
	validarNumero(b)
	validarNumero(c)
	validarNumero(d)
	validarNumero(sumarNumeros(5, 3))
	validarNumero(multiplicarNumeros(5, 3))

	fmt.Println()

	fmt.Println(operacionAritmetica(6, 2, suma))
	fmt.Println(operacionAritmetica(6, 2, resta))
	fmt.Println(operacionAritmetica(6, 2, multip))
	fmt.Println(operacionAritmetica(6, 2, "/"))

	validarNumero(sumaEllipsis(a))
	validarNumero(sumaEllipsis(a, b))
	validarNumero(sumaEllipsis(a, b, c))
	validarNumero(sumaEllipsis(a, b, c, d))

	fmt.Println("\nOperaciones: ")
	fmt.Println(operaciones(26, 36))

	s, r, m, e := operaciones2(26, 36)

	fmt.Println("Suma:\t\t", s)
	fmt.Println("Resta:\t\t", r)
	fmt.Println("Multiplicacion:\t", m)
	fmt.Println("Division:\t", e)

	res, err := division(2, 4)

	if err != nil {
		// Si hubo error
		fmt.Println("Hubo error")
	} else {
		// Si termino correctamente
		fmt.Println(res)
	}

	oper := operacionAritmetica2(divis)
	re := oper(2, 5)
	fmt.Println(re)

}

func operaciones2(valor1, valor2 float64) (suma float64, resta float64, multip float64, divis float64) {
	suma = valor1 + valor2
	resta = valor1 - valor2
	multip = valor1 * valor2

	if valor2 != 0 {
		divis = valor1 / valor2
	}

	return
}

func division(dividendo, divisor float64) (float64, error) {

	if divisor == 0 {
		return 0, errors.New("El divisor no puede ser cero")
	}

	return dividendo / divisor, nil
}

func validarNumero(numero int) {
	fmt.Printf("\nEl numero %d es ", numero)
	switch {
	case numero < 0:
		fmt.Printf("negativo")

	case numero > 0:
		fmt.Printf("positivo")

	default:
		fmt.Printf("cero")
	}
}
func opSuma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func opResta(valor1, valor2 float64) float64 {
	return valor1 - valor2
}

func opMultip(valor1, valor2 float64) float64 {
	return valor1 * valor2
}

func opDivis(valor1, valor2 float64) float64 {
	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2
}

func operacionAritmetica2(operador string) func(valor1, valor2 float64) float64 {
	switch operador {
	case suma:
		return opSuma
	case resta:
		return opResta
	case multip:
		return opMultip
	case divis:
		return opDivis
	}

	return nil
}
