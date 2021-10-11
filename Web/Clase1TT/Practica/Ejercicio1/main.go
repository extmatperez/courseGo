package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*Ejercicio 1 - Estructura un JSON
Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.
Los productos varían por id, nombre, color, precio, stock, código (alfanumérico),
publicado (si-no), fecha de creación.


Los usuarios varían por id, nombre, apellido, email, edad, altura, activo (si-no),
fecha de creación.
Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor (string), fecha de transacción.
Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el tema elegido,
ej: products.json.
Dentro del mismo escribí un JSON que permita tener un array de productos,
usuarios o transacciones con todas sus variantes.
*/

type Usuario struct {
	ID              int     `json:"id"`
	Nombre          string  `json:"nombre"`
	Apellido        string  `json:"apellido"`
	Email           string  `json:"email"`
	Edad            int     `json:"edad"`
	Altura          float64 `json:"altura"`
	Activo          bool    `json:"activo"`
	FechaDeCreacion string  `json:"fechaDeCreacion"`
}

func Funcion(ctxt *gin.Context) {

	// ctxt.String(200, "¡Bienvenido a la Empresa Gophers!")
	ctxt.JSON(200, gin.H{
		"message": "Hello World!",
	})

}
func Funcion2(ctxt *gin.Context) {

	// ctxt.String(200, "¡Bienvenido a la Empresa Gophers!")
	ctxt.JSON(200, gin.H{
		"message": ctxt.Param("nombre"),
	})

}
func BuscarUsuarioId(ctxt *gin.Context) {

	data, err := os.ReadFile("./usuarios.json")
	if err != nil {
		fmt.Printf("Error reading")
		ctxt.JSON(404, gin.H{
			"message": "Error reading",
		})
	} else {
		var usuarios []Usuario
		json.Unmarshal(data, &usuarios)
		//ctxt.JSON(200, usuarios)

		ID, _ := strconv.ParseInt(ctxt.Param("id"), 10, 64)
		for _, v := range usuarios {
			if int64(v.ID) == ID {
				ctxt.JSON(200, v)
				ID = -999
				break
			}

		}
		if ID != -999 {
			ctxt.JSON(404, gin.H{
				"message": "No se encontro el usuario buscado",
			})
		}
		// ctxt.String(200, string(data))
	}

}

func GetAll(ctxt *gin.Context) {

	data, err := os.ReadFile("./usuarios.json")
	if err != nil {
		fmt.Printf("Error reading")
		ctxt.JSON(404, gin.H{
			"message": "Error reading",
		})
	} else {
		ctxt.String(200, string(data))
	}

}

func GetAllFiltrado(ctxt *gin.Context) {

	data, err := os.ReadFile("./usuarios.json")
	if err != nil {
		fmt.Printf("Error reading")
		ctxt.JSON(404, gin.H{
			"message": "Error reading",
		})
	} else {
		var usuarios []Usuario
		var usuariosFiltrados []Usuario
		json.Unmarshal(data, &usuarios)
		//ctxt.JSON(200, usuarios)

		var edad int64
		edad, _ = strconv.ParseInt(ctxt.Query("nombre"), 10, 64)
		for _, v := range usuarios {
			if int64(v.Edad) <= edad {
				usuariosFiltrados = append(usuariosFiltrados, v)
			}

		}
		ctxt.JSON(200, usuariosFiltrados)
		// ctxt.String(200, string(data))
	}

}

func main() {
	// Crea un router con gin
	router := gin.Default()

	// Captura la solicitud GET “/hello-world”
	router.GET("/hola", Funcion)
	router.GET("/hola:nombre", Funcion2)
	router.GET("/usuarios/:id", BuscarUsuarioId)
	router.GET("/usuarios", GetAll)
	router.GET("/usuarios/filtrado", GetAllFiltrado)

	router.Run() // Corremos nuestro servidor sobre el puerto 8080
	// router.Run(":23665") // Corremos nuestro servidor sobre el puerto 23665
}
