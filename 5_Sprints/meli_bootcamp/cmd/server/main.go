package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, _ := sql.Open("sqlite3", "./meli.db")
	router := gin.Default()

	// Codigo de ayuda
	buyerRepository := buyer.NewRepository(db)
	buyerService := buyer.NewService(buyerRepository)
	warehouseHandler := handler.NewWarehouse(buyerService)
	warehousesRoutes := router.Group("/api/v1/warehouses")
	{
		warehousesRoutes.GET("/", warehouseHandler.GetAll())
		warehousesRoutes.GET("/:id", warehouseHandler.Get())
		warehousesRoutes.POST("/", warehouseHandler.Store())
		warehousesRoutes.PATCH("/:id", warehouseHandler.Update())
		warehousesRoutes.DELETE("/:id", warehouseHandler.Delete())
	}

	router.Run()
}
