package main

import (
	usuarios "github.com/extmatperez/courseGo/Web/Clase3TT/Practica/internal/usuarios"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := usuarios.NewRepository()
	serv := usuarios.NewService(repo)
	//cont := handler.NewUser(serv)

	r := gin.Default()
	pr := r.Group("/users")
	pr.POST("/", cont.Store())
	pr.GET("/", cont.GetAll())
	pr.GET("/load", cont.LoadFile())
	r.Run()
}
