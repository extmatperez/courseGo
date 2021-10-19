package handler

import (
	"github.com/extmatperez/courseGo/5_Sprints/meli_bootcamp/internal/buyer"
	"github.com/extmatperez/courseGo/5_Sprints/meli_bootcamp/internal/domain"
	"github.com/extmatperez/courseGo/5_Sprints/meli_bootcamp/pkg/web"
	"github.com/gin-gonic/gin"
)

type Buyer struct {
	buyerService buyer.Service
}

func NewBuyer(b buyer.Service) *Buyer {
	return &Buyer{
		buyerService: b,
	}
}

func (b *Buyer) Get() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		cardNumberID := c.Param("id")

		buyer, err := b.buyerService.Get(cardNumberID)

		if err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		} else {
			response := response{buyer}
			c.JSON(200, response)
		}

	}
}

func (b *Buyer) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Buyer `json:"data"`
	}

	return func(c *gin.Context) {
		lista, err := b.buyerService.GetTodos()
		if err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		} else {
			response := response{lista}
			c.JSON(200, response)
		}
	}
}

func (b *Buyer) Store() gin.HandlerFunc {
	type request struct {
		CardNumberID string `json:"card_number_id"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			// if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewError(401, "Token invalido"))

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
		p, err := b.buyerService.Crear(req.CardNumberID, req.FirstName, req.LastName)
		if err != nil {
			ctx.JSON(404, web.NewError(404, err.Error()))
			return
		}
		response := response{p}
		ctx.JSON(200, response)

	}
}

func (b *Buyer) Update() gin.HandlerFunc {
	type request struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		cardNumberID := ctx.Param("id")
		if token != "123456" {
			// if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewError(401, "Token invalido"))

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

		p, err := b.buyerService.Actualizar(cardNumberID, req.FirstName, req.LastName)
		if err != nil {
			ctx.JSON(404, web.NewError(404, err.Error()))
			return
		}
		response := response{p}
		ctx.JSON(200, response)

	}
}

func (b *Buyer) Delete() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		cardNumberID := ctx.Param("id")
		if token != "123456" {
			// if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewError(401, "Token invalido"))

			return
		}

		err := b.buyerService.Eliminar(cardNumberID)
		if err != nil {
			ctx.JSON(404, web.NewError(404, err.Error()))
			return
		}
		response := response{gin.H{"ok": "La eliminación fue exitosa"}}
		ctx.JSON(200, response)
	}
}
