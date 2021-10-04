package main

import "fmt"

/*Ejercicio 2 - Clima

Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de distintos lugares.
Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
Imprime los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?
*/

func main() {

	var temperatura float32 = 16.5
	var humedad float32 = 78.55
	var presion float32 = 925.95

	fmt.Printf("La temperatura es %.2f con una humedad del %.2f a una presión atmosférica del %.2f", temperatura, humedad, presion)

	//Las tres de tipo flotante
}
