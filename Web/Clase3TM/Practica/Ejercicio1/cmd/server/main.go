package main

import (
	"github.com/extmatperez/courseGo/Web/Clase3TM/Practica/Ejercicio1/cmd/server/handler"
	usuarios "github.com/extmatperez/courseGo/Web/Clase3TM/Practica/Ejercicio1/internal/usuarios"
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
	pr.PUT("/modificar/:id", cont.Modificar())
	pr.PATCH("/parchar/:id", cont.Patch())
	pr.DELETE("/eliminar/:id", cont.Eliminar())

	r.Run()
}
