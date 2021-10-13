package main

import (
	usuarios "github.com/extmatperez/courseGo/Web/Clase2TT/Practica/Ejercicio1/internal/usuarios"
)

func main() {
	repo := usuarios.NewRepository()
	serv := usuarios.NewService(repo)

}
