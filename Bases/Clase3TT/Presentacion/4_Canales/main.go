package main

import (
	"fmt"
	"time"
)

func proceso(i int, c chan int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Terminó")
	c <- i
}

func main() {
	c := make(chan int)
	go proceso(1, c)
	fmt.Println("Termino el programa")
	// <-c                                              // recibimos el valor del canal
	// variable := <-c // recibimos y lo asignamos a una variable
	fmt.Println("Termino el programa en ", <-c) // recibimos y lo imprimimos
	// fmt.Println("Termino el programa en ", variable) // recibimos y lo imprimimos

	// c := make(chan int)

	for i := 0; i < 10; i++ {
		go proceso(i, c)
	}

	<-c
	fmt.Println("Termino el programa en ", <-c)
	for i := 0; i < 8; i++ {
		fmt.Println("Termino el programa en ", <-c)
	}

}
