package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./archivo.txt")
	fmt.Println((data))
	fmt.Println(string(data))

	d1 := []byte("hello, gophers!")
	err := os.WriteFile("./myFile2.txt", d1, 0666)
	if err != nil {
		fmt.Println("Error de escritura")
	} else {
		fmt.Println("Escritura correcta")
	}
}
