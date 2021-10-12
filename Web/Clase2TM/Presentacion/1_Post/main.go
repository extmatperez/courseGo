package main

import "github.com/gin-gonic/gin"

type request struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

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

func main() {

	r := gin.Default()
	r.POST("/productos", FuncionPost)
	r.Run()

}
