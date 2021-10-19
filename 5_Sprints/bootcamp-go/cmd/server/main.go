package main

import (
	"github.com/extmatperez/meli_bootcamp10_sprints/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/buyer"
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/employee"
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/product"
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/section"
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/seller"
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/warehouse"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, _ := sql.Open("sqlite3", "./meli.db")
	router := gin.Default()

	// EMPLOYEE routes
	employeeRepository := employee.NewRepository(db)
	employeeService := employee.NewService(employeeRepository)
	employeeHandler := handler.NewEmployee(employeeService)

	employeeRoutes := router.Group("/api/v1/employees")
	{
		employeeRoutes.GET("/", employeeHandler.GetAll())
		employeeRoutes.GET("/:card_number_id", employeeHandler.Get())
		employeeRoutes.POST("/", employeeHandler.Store())
		employeeRoutes.PATCH("/:card_number_id", employeeHandler.Update())
		employeeRoutes.DELETE("/:card_number_id", employeeHandler.Delete())

	}

	// SELLER routes
	sellerRepository := seller.NewRepository(db)
	sellerService := seller.NewService(sellerRepository)
	sellerHandler := handler.NewSeller(*sellerService)

	sellersRoutes := router.Group("/api/v1/sellers")
	{
		sellersRoutes.GET("/", sellerHandler.GetAll())
		sellersRoutes.GET("/:cid", sellerHandler.Get())
		sellersRoutes.POST("/", sellerHandler.Store())
		sellersRoutes.PATCH("/:id", sellerHandler.Update())
		sellersRoutes.DELETE("/:id", sellerHandler.Delete())
	}

	// SECTION routes
	sectionRepository := section.NewRepository(db)
	sectionService := section.NewService(sectionRepository)
	sectionHandler := handler.NewSection(sectionService)

	sectionRoutes := router.Group("/api/v1/sections")
	{
		sectionRoutes.GET("/", sectionHandler.GetAll())
		sectionRoutes.GET("/:id", sectionHandler.Get())
		sectionRoutes.POST("/", sectionHandler.Store())
		sectionRoutes.PATCH("/:id", sectionHandler.Update())
		sectionRoutes.DELETE("/:id", sectionHandler.Delete())
	}

	// BUYER routes
	buyerRepository := buyer.NewRepository(db)
	buyerService := buyer.NewService(buyerRepository)
	buyerHandler := handler.NewBuyer(buyerService)

	buyersRoutes := router.Group("/api/v1/buyers")
	{
		buyersRoutes.POST("/", buyerHandler.Store())
		buyersRoutes.GET("/", buyerHandler.GetAll())
		buyersRoutes.GET("/:card_number_id", buyerHandler.Get())
		buyersRoutes.DELETE("/:card_number_id", buyerHandler.Delete())
		buyersRoutes.PATCH("/:card_number_id", buyerHandler.Update())
	}

	// WAREHOUSES Routes
	warehouseRepository := warehouse.NewRepository(db)
	warehouseService := warehouse.NewService(warehouseRepository)
	warehouseHandler := handler.NewWarehouse(*warehouseService)

	warehousesRoutes := router.Group("/api/v1/warehouses")
	{
		warehousesRoutes.GET("/", warehouseHandler.GetAll())
		warehousesRoutes.GET("/:id", warehouseHandler.Get())
		warehousesRoutes.POST("/", warehouseHandler.Store())
		warehousesRoutes.PATCH("/:id", warehouseHandler.Update())
		warehousesRoutes.DELETE("/:id", warehouseHandler.Delete())

	}

	// PRODUCT Routes
	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProduct(productService)
	productRoutes := router.Group("/api/v1/products")
	{
		productRoutes.GET("/", productHandler.GetAll())
		productRoutes.GET("/:id", productHandler.Get())
		productRoutes.POST("/", productHandler.Store())
		productRoutes.PATCH("/:id", productHandler.Update())
		productRoutes.DELETE("/:id", productHandler.Delete())

	}

	router.Run()
}
