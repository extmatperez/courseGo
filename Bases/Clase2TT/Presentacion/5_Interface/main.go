package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}

func (c circle) perim() float64 {
	return math.Pi * 2 * c.radius
}

func details(g geometry) {
	fmt.Printf("%T", g)
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func newCircle(values float64) circle {
	return circle{values}
}

func newRect(w, h float64) rect {
	return rect{w, h}
}

const (
	rectType   = "RECT"
	circleType = "CIRCLE"
)

func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case rectType:
		return rect{width: values[0], height: values[1]}
	case circleType:
		return circle{radius: values[0]}
	}
	return nil
}

func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	ci2 := newCircle(1)
	re2 := newRect(2, 3)
	details(r)
	details(c)
	details(ci2)
	details(re2)

	r1 := newGeometry(rectType, 2, 3)
	fmt.Println(r1.area())
	fmt.Println(r1.perim())
	c1 := newGeometry(circleType, 1)
	fmt.Println(c1.area())
	fmt.Println(c1.perim())

	details(r1)
	details(c1)

	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	fmt.Println(i)

}
