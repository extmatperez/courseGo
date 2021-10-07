package main

import (
	"errors"
	"fmt"
	"time"
)

// sólo se requiere crear un tipo que implemente el método Error()
type myCustomError struct {
	status int
	msg    string
}

//hacemos que nuestro tipo struct implemente el método Error()
func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

func myCustomErrorTest(status int) (int, error) {
	if status >= 300 {
		return 400, &myCustomError{
			status: status,
			msg:    "algo salió mal",
		}
	}
	return 200, nil
}

func main() {
	status, err := myCustomErrorTest(350) //llamamos a nuestra func
	if err != nil {                       //hacemos una validación del valor de err
		fmt.Println(err) //si err no es nil, imprimimos el error y...
		//os.Exit(1)       //utilizamos este método para salir del programa
	}
	fmt.Printf("Status %d, Funciona!\n", status)

	statusCode := 404
	if statusCode > 399 {
		fmt.Println(errors.New("la petición ha fallado."))
		//return
	}
	fmt.Println("El programa finalizó correctamente.")

	err = fmt.Errorf("momento del error: %v", time.Now())
	fmt.Println("error ocurrido: ", err)
}
