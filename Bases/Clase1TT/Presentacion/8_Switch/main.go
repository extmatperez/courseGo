package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	var t int = today.Day()
	switch t {
	case 5, 10, 15:
		fmt.Println("Limpia tu casa.")
	case 25, 26, 27:
		fmt.Println("Comprar comida.")
		fallthrough
	case 29:
		fmt.Println("Comemos ñoquis")
	default:
		fmt.Println("No hay información disponible para ese día.")
	}

	day := "domingo"

	switch day {
	case "lunes", "martes", "miércoles", "jueves", "viernes":
		fmt.Printf("%s es un día de la semana\n", day)
	default:
		fmt.Printf("%s es un día del fin de la semana\n", day)
	}

	switch day := "domingo"; day {
	case "lunes", "martes", "miércoles", "jueves", "viernes":
		fmt.Printf("%s es un día de la semana\n", day)
	default:
		fmt.Printf("%s es un día del fin de la semana\n", day)
	}

	var edad uint8 = 18
	switch {
	case edad >= 150:
		fmt.Println("¿Eres inmortal?")
	case edad >= 18:
		fmt.Println("Eres mayor de edad")
	default:
		fmt.Println("Eres menor de edad")
	}

}
