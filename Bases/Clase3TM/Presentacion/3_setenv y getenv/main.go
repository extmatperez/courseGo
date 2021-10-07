package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("NOMBRES", "gopher")

	value := os.Getenv("NOMBRES")
	fmt.Printf(value)
	// os.Exit(2)

	files, err := os.ReadDir(".")
	if err == nil {
		for _, v := range files {
			fmt.Println(v.Name())
			fmt.Println(v.Info())
			fmt.Println(v.Type())
		}

	}
}
