package main

import "fmt"

func main() {
	const nombre, edad = "Kim", 22
	fmt.Print(nombre, " tiene ", edad, " años de edad.\n")
	fmt.Println(nombre, " tiene ", edad, " años de edad.")
	res := fmt.Sprint(nombre, " tiene ", edad, " años de edad.\n")
	fmt.Print(res)

	res = fmt.Sprintln(nombre, " tiene ", edad, " años de edad.")
	fmt.Print(res)
}
