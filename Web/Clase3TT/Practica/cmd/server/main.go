package main

import (
	"github.com/extmatperez/courseGo/Web/Clase2TT/Practica/Ejercicio1/cmd/server/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := usuarios.NewRepository()
	serv := usuarios.NewService(repo)
	cont := handler.NewUser(serv)

	r := gin.Default()
	pr := r.Group("/users")
	pr.POST("/", cont.Store())
	pr.GET("/", cont.GetAll())
	pr.GET("/load", cont.LoadFile())
	r.Run()
}
