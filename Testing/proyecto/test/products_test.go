package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/extmatperez/w2GoPrueba/GoTesting/Clase2TT/proyecto/cmd/server/handler"
	personas "github.com/extmatperez/w2GoPrueba/GoTesting/Clase2TT/proyecto/internal/personas"
	"github.com/extmatperez/w2GoPrueba/GoTesting/Clase2TT/proyecto/pkg/store"
	"github.com/extmatperez/w2GoPrueba/GoTesting/Clase2TT/proyecto/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	router := gin.Default()

	db := store.New(store.FileType, "./personasSalidaTest.json")
	repo := personas.NewRepository(db)
	service := personas.NewService(repo)
	controller := handler.NewPersona(service)

	router.GET("/personas/get", controller.GetAll())
	router.POST("/personas/add", controller.Store())
	router.PUT("/personas/:id", controller.Update())
	router.PATCH("/personas/:id", controller.UpdateNombre())
	router.DELETE("/personas/:id", controller.Delete())
	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetProduct_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/personas/get", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	var objRes web.Response
	assert.Equal(t, 200, rr.Code)

	//fmt.Println(string(rr.Body.Bytes()))
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// fmt.Println(string(rr.Body.Bytes()))
	assert.Nil(t, err)
}
