package handler

import (
	usuarios "github.com/extmatperez/courseGo/Web/Clase3TT/Practica/internal/usuarios"
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

func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
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
		if token != "123456" {
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

func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}
