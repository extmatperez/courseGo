package main

import (
	"github.com/extmatperez/courseGo/Web/Clase3TT/Presentacion/MVC/cmd/server/handler"
	products "github.com/extmatperez/courseGo/Web/Clase3TT/Presentacion/MVC/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()
}
