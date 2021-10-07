package main

import "fmt"

func main() {
	var p1 *int
	var p2 = new(int)
	v := 5.5
	p3 := &v

	fmt.Println(p1)
	fmt.Println(*p2)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(*p3)

	var b int = 19
	fmt.Println("La dirección de memoria de v es: ", &b, " y el valor es ", b)
}
