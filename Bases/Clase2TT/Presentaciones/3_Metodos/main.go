package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	Radio float64 `json:"radio"`
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(float64(c.Radio), 2)
}

func (c Circulo) Perimetro() float64 {
	return math.Pi * 2 * c.Radio
}

func (c *Circulo) setRadio(nuevoRadio float64) {
	c.Radio = nuevoRadio
}

func main() {
	miCirculo := Circulo{2.5}

	fmt.Printf("Area: %.2f, Perimetro: %.2f\n", miCirculo.Area(), miCirculo.Perimetro())

	fmt.Println(miCirculo)
	miCirculo.setRadio(23)
	fmt.Println(miCirculo)
}
