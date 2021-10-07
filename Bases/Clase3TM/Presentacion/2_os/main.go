package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "./archivos/archivo.txt"
	f, err := os.Stat(filename)
	if err == nil {
		fmt.Println("Es un directorio: ", f.IsDir())
		fmt.Println("Nombre del archivo/directorio: ", f.Name())
		fmt.Println("Tamaño del archivo en Bytes: ", f.Size())
		fmt.Println("Fecha y Hora del archivo: ", f.ModTime())
		fmt.Println("Permisos del archivo", f.Mode())
	} else {
		fmt.Println("El archivo no existe")
	}
}
