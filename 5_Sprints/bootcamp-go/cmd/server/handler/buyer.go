package handler

import (
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/buyer"
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/pkg/web"
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

// Receives a card_number_id as string and returns the buyer associated
func (b *Buyer) Get() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {

		cardNumberID := c.Param("card_number_id")

		if buyer, err := b.buyerService.Get(cardNumberID); err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		} else {
			response := response{Data: buyer}
			c.JSON(200, response)
		}
	}
}

// Returns all the existing buyers
func (b *Buyer) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Buyer `json:"data"`
	}

	return func(c *gin.Context) {

		if buyersList, err := b.buyerService.GetAll(); err != nil {
			c.JSON(500, web.NewError(500, err.Error()))
			return
		} else {
			response := response{Data: buyersList}
			c.JSON(200, response)
		}
	}
}

// Receives a card_number_id, first_name and last_name as strings, creates a buyer and returns it
func (b *Buyer) Store() gin.HandlerFunc {
	type request struct {
		CardNumberID string `json:"card_number_id" binding:"required"`
		FirstName    string `json:"first_name" binding:"required"`
		LastName     string `json:"last_name" binding:"required"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		var req request

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(422, web.NewError(422, err.Error()))
			return
		}

		newBuyer, err := b.buyerService.Create(req.CardNumberID, req.FirstName, req.LastName)
		if err != nil {
			c.JSON(409, web.NewError(409, err.Error()))
			return
		}

		response := response{Data: newBuyer}
		c.JSON(201, response)
	}
}

// Receives a card_number_id, first_name and last_name as strings, updates buyer for given card_number_id and returns it updated
func (b *Buyer) Update() gin.HandlerFunc {
	type request struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		var req request

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(422, web.NewError(422, err.Error()))
			return
		}

		cardNumberID := c.Param("card_number_id")

		if buyer, err := b.buyerService.Update(cardNumberID, req.FirstName, req.LastName); err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		} else {

			c.JSON(200, response{Data: buyer})
		}
	}
}

// Receives a card_number_id and deletes the buyer associated with it
func (b *Buyer) Delete() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {

		cardNumberID := c.Param("card_number_id")

		if err := b.buyerService.Delete(cardNumberID); err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		}

		c.JSON(204, response{})
	}
}
