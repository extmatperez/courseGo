package main

import "fmt"

type Vehiculo struct {
	km     float64
	tiempo float64
}

type Auto struct {
	v Vehiculo
}

type Moto struct {
	v Vehiculo
}

//agregaremos un método que reciba tiempo en minutos y se encargue de realizar el cálculo de distancia en base a 100km/h
func (a *Auto) Correr(minutos int) {
	a.v.tiempo = float64(minutos) / 60
	a.v.km = a.v.tiempo * 100
}

//Agregamos el método Correr que recibe el tiempo en minutos y hace el cálculo en base a 80km/h

func (m *Moto) Correr(minutos int) {
	m.v.tiempo = float64(minutos) / 60
	m.v.km = m.v.tiempo * 80
}

func (v Vehiculo) detalle() {
	fmt.Printf("km:\t%f\ntiempo:\t%f\n", v.km, v.tiempo)
}

//y el método detalle que llame al método de la clase padre
func (a *Auto) Detalle() {
	fmt.Println("\nV:\tAuto")
	a.v.detalle()
}

func (m *Moto) Detalle() {
	fmt.Println("\nV:\tMoto")
	m.v.detalle()
}

func main() {
	auto := Auto{}
	auto.Correr(360)
	auto.Detalle()

	moto := Moto{}
	moto.Correr(360)
	moto.Detalle()
}
