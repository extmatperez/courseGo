package main

import (
	"github.com/extmatperez/courseGo/Web/Clase3TT/Practica/cmd/server/handler"
	usuarios "github.com/extmatperez/courseGo/Web/Clase3TT/Practica/internal/usuarios"
	"github.com/extmatperez/courseGo/Web/Clase3TT/Practica/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./usuariosSalida.json")
	repo := usuarios.NewRepository(db)
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
