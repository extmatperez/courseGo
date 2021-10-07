package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	_, err := ioutil.ReadFile("customers.txt")

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Ejecución finalizada")
	}()

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	// defer func() {
	// 	fmt.Println("Ejecución finalizada")
	// }()
	// fmt.Println("Esto va primero")
	// if err != nil {
	// 	panic("el archivo indicado no fue encontrado o está dañado")
	// }

}
