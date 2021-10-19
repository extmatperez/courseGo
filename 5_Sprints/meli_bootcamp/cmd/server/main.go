package main

import (
	"database/sql"

	"github.com/extmatperez/courseGo/5_Sprints/meli_bootcamp/cmd/server/handler"
	"github.com/extmatperez/courseGo/5_Sprints/meli_bootcamp/internal/buyer"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// db, _ := sql.Open("sqlite3", "./meli.db")
	db, _ := sql.Open("mysql", "root:root@/meli")
	router := gin.Default()

	// Codigo de ayuda
	buyerRepository := buyer.NewRepository(db)
	buyerService := buyer.NewService(buyerRepository)
	buyerHandler := handler.NewBuyer(buyerService)

	buyerRoutes := router.Group("/api/v1/buyer")
	{
		buyerRoutes.GET("/", buyerHandler.GetAll())
		buyerRoutes.GET("/:id", buyerHandler.Get())
		buyerRoutes.POST("/", buyerHandler.Store())
		buyerRoutes.PATCH("/:id", buyerHandler.Update())
		buyerRoutes.DELETE("/:id", buyerHandler.Delete())
	}

	router.Run()
}
