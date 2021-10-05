package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Hola Mundo!")
	var numero int = 5
	numero2 := 25
	fmt.Println(numero)
	fmt.Println(numero2)

	numero = -9
	fmt.Println(numero)

	nombre1, nombre2 := "Hola   df ", 10.5
	nombre3 := true

	var (
		product  = "Course"
		quantity = 25
		price    = 40.50
		inStock  = true
	)
	fmt.Printf("%v => %T => %d bytes\n", nombre1, nombre1, unsafe.Sizeof(nombre1))
	fmt.Printf("%v => %T => %d bytes\n", nombre2, nombre2, unsafe.Sizeof(nombre2))
	fmt.Printf("%v => %T => %d bytes\n", product, product, unsafe.Sizeof(product))
	fmt.Println(nombre2)
	fmt.Println(nombre3)
	fmt.Println(product)
	fmt.Println(quantity)
	fmt.Println(price)
	fmt.Println(inStock)

	const (
		PRODUCT  = "Course"
		QUANTITY = 20
		PRICE    = 40.50
	)

	fmt.Println(PRODUCT)
	fmt.Println(QUANTITY)
	fmt.Println(PRICE)

	var year = 2021
	fmt.Printf("%v => %T => %d bytes\n", PRODUCT, PRODUCT, unsafe.Sizeof(year))

	var fname, lname string = "John", "Doe"

	fmt.Println(fname, lname)

	var algo1 = "A"
	fmt.Printf("%v => %T => %d bytes\n", algo1, algo1, unsafe.Sizeof(algo1))

	var algo2 byte = 5
	fmt.Printf("%v => %T => %d bytes\n", algo2, algo2, unsafe.Sizeof(algo2))

	var algo3 = "Aa"
	fmt.Printf("%v => %T => %d bytes\n", algo3, algo3, unsafe.Sizeof(algo3))

}
