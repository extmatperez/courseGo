package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/extmatperez/w2GoPrueba/GoTesting/Clase2TT/proyecto/pkg/store"
	productos "github.com/extmatperez/w2GoPrueba/GoTesting/Ejemplo/proyecto/internal/productos"
	"github.com/extmatperez/w2GoPrueba/GoWeb/Clase4TM/proyecto/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

func createServer() *gin.Engine {
	router := gin.Default()
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./ProductosSalidaTest.json")
	repo := productos.NewRepository(db)
	service := productos.NewService(repo)
	controller := NewProducto(service)

	router.Use(TokenAuthMiddleware())

	router.GET("/Productos/get", controller.GetAll())
	router.POST("/Productos/add", controller.Store())
	router.PUT("/Productos/:id", controller.Update())
	router.DELETE("/Productos/:id", controller.Delete())

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetProductos(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest(http.MethodGet, "/Productos/get", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)

	assert.Equal(t, 200, respuesta.Code)
	assert.Nil(t, err)
}

func Test_StoreProductos(t *testing.T) {
	router := createServer()

	nuevaProducto := productos.Producto{Nombre: "Horno",
		Precio: 26354}

	dataNueva, _ := json.Marshal(nuevaProducto)
	fmt.Println(string(dataNueva))
	req, rr := createRequestTest(http.MethodPost, "/Productos/add", string(dataNueva))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta productos.Producto

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	assert.Equal(t, "Horno", respuesta.Nombre)
	assert.Nil(t, err)

	// delete := fmt.Sprintf("/Productos/%d", respuesta.ID)
	// req, rr = createRequestTest(http.MethodDelete, delete, "")

	// router.ServeHTTP(rr, req)

	// assert.Equal(t, 200, rr.Code)

}
