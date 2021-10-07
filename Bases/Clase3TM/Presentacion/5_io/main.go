package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	data, _ := os.ReadFile("./archivo.txt")
	// fmt.Println((data))
	// fmt.Println(string(data))
	r := strings.NewReader(string(data))
	v, err := io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Printf("Error")
	}
	fmt.Println(v)

	r = strings.NewReader(string(data))

	b, err := io.ReadAll(r)
	fmt.Println(string(b), err)

	io.WriteString(os.Stdout, "Hello world!")

	// fmt.Println(string(v), err)
}
