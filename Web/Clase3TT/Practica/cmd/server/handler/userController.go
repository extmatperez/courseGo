package handler

import (
	"fmt"
	"os"
	"strconv"

	usuarios "github.com/extmatperez/courseGo/Web/Clase3TT/Practica/internal/usuarios"
	"github.com/extmatperez/courseGo/Web/Clase3TT/Practica/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre          string  `json:"nombre"`
	Apellido        string  `json:"apellido"`
	Email           string  `json:"email"`
	Edad            int     `json:"edad"`
	Altura          float64 `json:"altura"`
	Activo          bool    `json:"activo"`
	FechaDeCreacion string  `json:"fechaDeCreacion"`
}

type User struct {
	service usuarios.Service
}

func NewUser(p usuarios.Service) *User {
	return &User{
		service: p,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(p) > 0 {
			ctx.JSON(200, p)
		} else {
			ctx.JSON(404, gin.H{"error": "No hay usuarios cargados"})
		}
	}
}

func (c *User) LoadFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.LoadFile()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(p) > 0 {
			ctx.JSON(200, p)
		} else {
			ctx.JSON(404, gin.H{"error": "No hay usuarios cargados"})
		}
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		// if os.Getenv("TOKEN") == "" {
		// 	fmt.Printf("Sin token")
		// } else {
		// 	fmt.Println(os.Getenv("TOKEN"))
		// }
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))

			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		//fmt.Println(req)
		p, err := c.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// ModificarUsuario godoc
// @Summary Modificar usuarios
// @Tags Usuarios
// @Description modificar usuarios
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param usuario body request true "Usuario a modificar"
// @Success 200 {object} web.Response
// @Router /usuarios [patch]
func (c *User) Modificar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		p, err := c.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}
func (c *User) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		p, err := c.service.UpdateNombre(int(id), req.Nombre)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *User) Eliminar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El producto %d ha sido eliminado", id), ""))
		// ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}
