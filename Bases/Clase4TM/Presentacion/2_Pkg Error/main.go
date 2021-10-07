package main

import (
	"errors"
	"fmt"
)

// definimos un type struct
type myError struct {
	msg string
	x   string
}

//hacemos que nuestro type struct implemente el método Error()
func (e *myError) Error() string {
	return fmt.Sprintf("ha ocurrido un error: %s, %s", e.msg, e.x)
}

// sólo se requiere crear un tipo que implemente el método Error()
type myCustomError struct {
	status int
	msg    string
}

//hacemos que nuestro tipo struct implemente el método Error()
func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

var err1 = errors.New("error número 1")

func x() error {
	return fmt.Errorf("información extra del error: %w", err1)
}

type errorTwo struct{}

func (e errorTwo) Error() string {
	return "error two happened"
}

func SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("nil name")
	}
	return fmt.Sprintf("Hola %s ", name), nil
}

func main() {
	e := &myError{"nuevo error", "404"}
	e2 := &myCustomError{205, "Error"}
	var err *myError

	isMyError := errors.As(e, &err) // compara los errores

	fmt.Println(isMyError)          //imprime true porque los errores coinciden
	isMyError = errors.As(e2, &err) // compara los errores

	fmt.Println(isMyError) //imprime true porque los errores coinciden

	e1 := x()
	coincidence := errors.Is(e1, err1)
	fmt.Println(coincidence) //imprime true

	fmt.Println(errors.Unwrap(e1)) //imprime e2

	e3 := errorTwo{}
	e4 := fmt.Errorf("e3: %w", e3)
	fmt.Println(errors.Unwrap(e4)) //imprime e2
	fmt.Println(errors.Unwrap(e3)) //imprime nil

	name := ""
	greeting, err5 := SayHello(name)
	if err5 != nil {
		fmt.Println("no se puede saludar", err5)
	}
	fmt.Println(greeting)
}
