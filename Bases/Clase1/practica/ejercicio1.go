package main

import "fmt"

/*Ejercicio 1 - Imprimí tu nombre
Crea una aplicación donde tengas como variable tu nombre y dirección.
Imprime en consola el valor de cada una de las variables.
*/

func main() {
	var (
		nombre    = "Matias"
		direccion = "San Francisco 64"
	)

	fmt.Printf("Mi nombre es: %s y vivo en: %s\n", nombre, direccion)
	fmt.Printf("Mi nombre es: %v y vivo en: %v", nombre, direccion)

}
