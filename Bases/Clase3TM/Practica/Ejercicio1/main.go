package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto,
con la información de productos comprados, separados por punto y coma.
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

type Producto struct {
	Id       int     `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad float64 `json:"cantidad"`
}

func main() {
	var misProductos []Producto

	p1 := Producto{1, 25.3, 25}
	p2 := Producto{2, 35.3, 12}
	p3 := Producto{3, 45.3, 36}
	p4 := Producto{4, 55.3, 96}
	p5 := Producto{5, 65.3, 12}

	misProductos = append(misProductos, p1, p2, p3, p4, p5)
	// fmt.Printf("%v\n", misProductos)
	salida := ""
	dato := ""
	for i := 0; i < len(misProductos)-1; i++ {
		str, _ := json.Marshal(misProductos[i])
		dato = string(str)
		salida += fmt.Sprint(dato, ";")
	}
	str, _ := json.Marshal(misProductos[len(misProductos)-1])
	dato = string(str)
	salida += fmt.Sprint(dato)
	// fmt.Printf("%v", salida)

	// r := strings.NewReader(string(salida))

	// b, _ := io.ReadAll(r)

	os.WriteFile("./Salida1.txt", []byte(salida), 0666)

	b, _ := json.Marshal(misProductos)

	os.WriteFile("./Salida2.txt", b, 0666)
	// producto := "{id: 1, nombre: shampoo, precio: $100, cantidad: 10};{id: 2, nombre: cepillo, precio: $50, cantidad: 12};{id: 2, nombre: jabon, precio: $75, cantidad: 21}"
	// archivo := []byte(producto)
	// err := ioutil.WriteFile("./Salida3.txt", archivo, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
