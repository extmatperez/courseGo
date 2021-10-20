package ordenamientos

import "fmt"

// Función que recibe dos enteros y retorna la suma resultante
func Sumar(num1, num2 int) int {
	return num1 + num2
}

// Función que recibe dos enteros y retorna la resta o diferencia resultante
func Restar(num1, num2 int) int {
	return num1 - num2
}

// Función que recibe dos enteros (numerador y denominador) y retorna la división resultante
func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("El denominador no puede ser 0")
	}
	return num / den, nil
}
