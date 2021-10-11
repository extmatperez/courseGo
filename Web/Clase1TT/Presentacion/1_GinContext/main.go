package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//Esta función nos permite ver la anatomía de un mensaje Request de una
func Ejemplo(context *gin.Context) {
	//El body, header y method están contenidos en el contexto de gin.
	contenido := context.Request.Body
	header := context.Request.Header
	metodo := context.Request.Method

	fmt.Println("¡He recibido algo!")
	fmt.Printf("\tMetodo: %s\n", metodo)
	fmt.Printf("\tContenido:\n")

	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	}
	fmt.Printf("\tCotenido:%s\n", contenido)
	fmt.Println("¡Yay!")
	context.String(200, "¡Lo recibí!") //Respondemos al cliente con 200 OK y un mensaje.
}

func main() {
	// Crea un router con gin
	router := gin.Default()

	// Captura la solicitud GET “/hello-world”
	router.GET("/", Ejemplo)

	gopher := router.Group("/gophers")
	{
		gopher.GET("/", Ejemplo)
		gopher.GET("/get", Ejemplo)
		gopher.GET("/info", Ejemplo)
	}

	router.Run() // Corremos nuestro servidor sobre el puerto 8080
	// router.Run(":23665") // Corremos nuestro servidor sobre el puerto 23665
}
