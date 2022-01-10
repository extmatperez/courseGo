package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
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

var misUsuarios []Usuario

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
		edad, _ = strconv.ParseInt(ctxt.Query("edad"), 10, 64)
		for _, v := range usuarios {
			if int64(v.Edad) <= edad {
				usuariosFiltrados = append(usuariosFiltrados, v)
			}

		}
		ctxt.JSON(200, usuariosFiltrados)
		// ctxt.String(200, string(data))
	}

}

func NuevoUsuario(c *gin.Context) {
	var req Usuario
	// c.Bind(&req)
	if err := c.Bind(&req); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return //Para que?
	}
	req.ID = len(misUsuarios) + 1
	c.JSON(200, req)
	misUsuarios = append(misUsuarios, req)

}

func validarUsuario(user Usuario, c *gin.Context) bool {
	gType := reflect.TypeOf(user)
	gValue := reflect.ValueOf(user)
	strNumFields := gType.NumField()
	var campos []string
	for i := 1; i < strNumFields; i++ {
		field := gType.Field(i)

		largo := gValue.Field(i).Interface()
		if largo == "" || largo == 0 || largo == 0.0 {
			campos = append(campos, field.Name)
			// fmt.Printf("\nEl campo %s es requerido", field.Name)
		}
	}
	if len(campos) == 0 {
		return true
	} else {
		if len(campos) == 1 {
			c.String(400, "El campo %s es requerido", campos)
		} else {
			c.String(400, "Los campos %s son requeridos", campos)
		}
		return false
	}
}

func NuevoUsuarioValidando(c *gin.Context) {

	if c.Request.Header.Get("token") != "mperez" {
		c.String(401, "No tiene permisos para realizar la petición solicitada")
		return
	}

	var req Usuario
	// c.Bind(&req)
	if err := c.Bind(&req); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return //Para que?
	}

	if validarUsuario(req, c) {
		req.ID = len(misUsuarios) + 1
		c.JSON(200, req)
		misUsuarios = append(misUsuarios, req)
	}

}

func GetAllUser(ctxt *gin.Context) {

	if len(misUsuarios) == 0 {
		ctxt.JSON(404, gin.H{"message": "No hay usuarios cargados aun"})
	} else {
		ctxt.JSON(200, misUsuarios)
	}

}

func main() {
	// Crea un router con gin
	router := gin.Default()
	// Captura la solicitud GET “/hello-world”
	router.GET("/hola", Funcion)
	router.GET("/hola/:nombre", Funcion2)
	group := router.Group("usuarios")
	group.GET("/:id", BuscarUsuarioId)
	group.GET("/", GetAll)
	group.GET("/filtrado", GetAllFiltrado)
	group.GET("/todos", GetAllUser)
	group.POST("/add", NuevoUsuario)
	group.POST("/addValid", NuevoUsuarioValidando)

	router.Run() // Corremos nuestro servidor sobre el puerto 8080
	// router.Run(":23665") // Corremos nuestlsro servidor sobre el puerto 23665
}



package main

import (
	"fmt"
	"reflect"
)

type Persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	personita := Persona{}
	fmt.Println(personita)
	var personita2 Persona
	fmt.Println(personita2)
	personita2.apellido = "Perez"
	fmt.Printf("personita2: %v\n", personita2)
	fmt.Println(reflect.TypeOf(personita2))
	fmt.Println(reflect.ValueOf(personita2))
	fmt.Println(reflect.TypeOf(personita2).Field(0).)
	fmt.Println(reflect.ValueOf(personita2).Field(0))
	fmt.Println(reflect.TypeOf(personita2).Field(1))
	fmt.Println(reflect.ValueOf(personita2).Field(1))
	fmt.Println(reflect.TypeOf(personita2).Field(2))
	fmt.Println(reflect.ValueOf(personita2).Field(2))
}
