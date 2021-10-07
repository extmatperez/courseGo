package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado,
para ello requiere que: se imprima por pantalla mostrando los valores tabulados
, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad),
 el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50

*/

type Producto struct {
	Id       int     `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad float64 `json:"cantidad"`
}

func main() {
	data, err := os.ReadFile("./Salida2.txt")
	if err != nil {
		fmt.Printf("Error reading")
		os.Exit(1)
	}
	fmt.Println((data))
	fmt.Println(string(data))

	var productos []Producto
	// productos := &[]Producto{}
	json.Unmarshal(data, &productos)

	fmt.Println(productos)
}
