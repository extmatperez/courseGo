package main

import (
	"fmt"
	"runtime"
	"time"
)

func proceso(i int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond) // time.Sleep(time.Second)
	fmt.Println(i, "-Terminó")
}

func main() {

	fmt.Printf("Esta compu tiene %d hilos", runtime.NumCPU())
	//Secuencial
	// for i := 0; i < 10; i++ {
	// 	proceso(i)
	// }
	// time.Sleep(5000 * time.Millisecond)
	// fmt.Println("Termino el programa")

	//Paralelo
	for i := 0; i < 10; i++ {
		go proceso(i)
	}
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Termino el programa")
}
