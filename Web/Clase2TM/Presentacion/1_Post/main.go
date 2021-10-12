package main

import "github.com/gin-gonic/gin"

type request struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

var products []request
var lastID int

func FuncionPost(c *gin.Context) {
	var req request
	// c.Bind(&req)
	if err := c.Bind(&req); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return //Para que?
	}
	req.ID = 4
	c.JSON(200, req)
}

func FuncionConHeader(c *gin.Context) {
	var req request
	// c.Bind(&req)
	if c.Request.Header.Get("token") == "123456" {
		if err := c.Bind(&req); err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
			return //Para que?
		}
		req.ID = 4
		c.JSON(200, req)
	} else {
		c.JSON(401, gin.H{"error": "Token incorrecto"})
	}
}

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.Bind(&req); err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		lastID++
		req.ID = lastID
		products = append(products, req)
		c.JSON(200, req)
	}
}

func Mostrar(c *gin.Context) {
	// c.Bind(&req)
	if len(products) == 0 {
		c.JSON(200, gin.H{"mensaje": "No hay productos cargados"})
		return //Para que?
	}
	c.JSON(200, products)
}
func main() {

	r := gin.Default()
	pr := r.Group("/productos")
	pr.POST("/", Guardar())
	pr.POST("/nuevo/", FuncionConHeader)
	pr.GET("/", Mostrar)
	r.Run()

}
