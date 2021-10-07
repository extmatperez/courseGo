package main

import (
	"fmt"
)

type Perro struct {
	Name string
}

func (s *Perro) Ladrar() {
	fmt.Println(s.Name, "hace guau guau")
}

func main() {
	// fmt.Println("Iniciando... ")
	// _, err := os.Open("main1.go")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Terminando...")

	// animals := []string{
	// 	"vaca",
	// 	"perro",
	// 	"halcon",
	// }
	// fmt.Println("solo vuela el: ", animals[len(animals)])

	// perrito := &Perro{"Moro"}
	// perrito.Ladrar()
	// perrito = nil
	// perrito.Ladrar()

	//aplicamos “defer” a la invocación de una función anónima
	defer func() {
		fmt.Println("Esta función se ejecuta a pesar de producirse panic")
	}()
	defer func() {
		fmt.Println("Esta función se ejecuta a pesar de producirse panic")
	}()
	func() {
		fmt.Println("Esta función se ejecuta a pesar de producirse panic")
	}()
	defer func() {
		fmt.Println("Esta función se ejecuta a pesar de producirse panic")
	}()
	//creamos un panic con un mensaje de que se produjo
	panic("se produjo panic!!!")

}

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {

// 	f := createFile("./defer.txt")
// 	defer closeFile(f)
// 	writeFile(f)
// }

// func createFile(p string) *os.File {
// 	fmt.Println("creating")
// 	f, err := os.Create(p)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return f
// }

// func writeFile(f *os.File) {
// 	fmt.Println("writing")
// 	fmt.Fprintln(f, "data")

// }

// func closeFile(f *os.File) {
// 	fmt.Println("closing")
// 	err := f.Close()

// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %v\n", err)
// 		os.Exit(1)
// 	}
// }
