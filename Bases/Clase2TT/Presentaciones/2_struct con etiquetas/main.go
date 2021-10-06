package main

import (
	"encoding/json"
	"fmt"
)

type Fecha struct {
	Dia  int //`json:"dia"`
	Mes  int //`json:"mes"`
	Anio int //`json:"año"`
}

type Persona struct {
	Nombre          string `json:"primer_nombre"`
	Edad            int    `json:"edad"`
	FechaNacimiento Fecha  //`json:"fecha_nacimiento"`
}

func main() {
	p1 := Persona{"Matias", 26, Fecha{17, 11, 1994}}

	miJson, err := json.Marshal(p1)

	fmt.Println(p1)
	//fmt.Println(miJson)
	fmt.Println(string(miJson))

	fmt.Println(err)

	fmt.Print(string(miJson))
	fmt.Printf("\n%v", string(miJson))

}
