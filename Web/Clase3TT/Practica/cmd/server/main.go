package main

import (
	"log"
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

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}

		c.Next()
	}
}

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


//router.Use(TokenAuthMiddleware())

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
