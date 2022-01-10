package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type x struct {
		A string
		B int
		C bool
	}

	z := x{
		A: "Hola Mundo",
		B: 5,
		C: true,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(z)

	fmt.Print(encoder.)
	fmt.Printf("%+v", z)

}
