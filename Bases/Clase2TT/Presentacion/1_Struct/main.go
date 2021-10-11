package main

import "fmt"

type Fecha struct {
	dia  int
	mes  int
	anio int
}

type Persona struct {
	nombre          string
	edad            int
	fechaNacimiento Fecha
}

func main() {
	p1 := Persona{"Matias", 26, Fecha{17, 11, 1994}}

	p2 := Persona{
		nombre:          "Matias2",
		edad:            26,
		fechaNacimiento: Fecha{17, 11, 1994},
	}

	var p3 Persona
	p3.nombre = "Matias"

	var p4 Persona
	p4.edad = 26

	// var p5 Persona{}
	// p5.nombre = "Matias"

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)

	p4 = p1
	p4.fechaNacimiento.dia = 16
	fmt.Println(p4)
}
