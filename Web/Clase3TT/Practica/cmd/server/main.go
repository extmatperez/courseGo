package main

import (
	"os"

	"github.com/extmatperez/courseGo/Web/Clase3TT/Practica/cmd/server/handler"
	"github.com/extmatperez/courseGo/Web/Clase3TT/Practica/docs"
	usuarios "github.com/extmatperez/courseGo/Web/Clase3TT/Practica/internal/usuarios"
	"github.com/extmatperez/courseGo/Web/Clase3TT/Practica/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Esto es un titulo
// @version 1.0
// @description Esto es una descripcion.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./usuariosSalida.json")
	repo := usuarios.NewRepository(db)
	serv := usuarios.NewService(repo)
	cont := handler.NewUser(serv)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/users")
	pr.POST("/", cont.Store())
	pr.GET("/", cont.GetAll())
	pr.GET("/load", cont.LoadFile())
	pr.PUT("/modificar/:id", cont.Modificar())
	pr.PATCH("/parchar/:id", cont.Patch())
	pr.DELETE("/eliminar/:id", cont.Eliminar())

	r.Run()
}
