package main

import "fmt"

/*Ejercicio 6 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a
imprimir la edad de Benjamin.

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario:
Saber cuántos de sus empleados son mayores a 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.
*/

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("La edad de Benjamin es: %d", employees["Benjamin"])

	var empleados = []string{}

	cantidad := 0
	var k string
	var v int
	for k, v = range employees {
		if v > 21 {
			cantidad++
			empleados = append(empleados, k)
		}
	}
	fmt.Printf("\nCantidad de empleados con mas de 21 años: %d ", cantidad)
	fmt.Println(empleados)

	//Agregar a Federico con 25 años

	employees["Federico"] = 25
	fmt.Println(employees)

	//Eliminar a Pedro

	delete(employees, "Pedro")
	fmt.Println(employees)

}
