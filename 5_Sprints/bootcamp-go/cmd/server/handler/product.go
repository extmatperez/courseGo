package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/product"
	"github.com/extmatperez/meli_bootcamp10_sprints/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	productService product.Service
}

func NewProduct(s product.Service) *Product {
	return &Product{
		productService: s,
	}
}

//Store: recibe un request para guardar un producto
func (p *Product) Store() gin.HandlerFunc {

	type request struct {
		Description    string  `json:"description" binding:"required"`
		ExpirationRate int     `json:"expiration_rate" binding:"required"`
		FreezingRate   int     `json:"freezing_rate" binding:"required"`
		Height         float64 `json:"height" binding:"required"`
		Length         float64 `json:"length" binding:"required"`
		Netweight      float64 `json:"netweight" binding:"required"`
		ProductCode    string  `json:"product_code" binding:"required"`
		RecomFreezTemp float64 `json:"recommended_freezing_temperature" binding:"required"`
		Width          float64 `json:"width"  binding:"required"`
		ProductTypeID  int     `json:"product_type_id" binding:"required" validate:"number"`
		SellerID       int     `json:"seller_id" binding:"required" validate:"number"`
	}

	type response struct {
		Data domain.Product `json:"data"`
	}

	return func(ctx *gin.Context) {

		var req request

		//si el JSON no contiene los campos necesarios se devuelve un 422
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		var productReq domain.Product
		productReq.Description = req.Description
		productReq.ExpirationRate = req.ExpirationRate
		productReq.FreezingRate = req.FreezingRate
		productReq.Height = req.Height
		productReq.Length = req.Length
		productReq.Netweight = req.Netweight
		productReq.ProductCode = req.ProductCode
		productReq.RecomFreezTemp = req.RecomFreezTemp
		productReq.Width = req.Width
		productReq.ProductTypeID = req.ProductTypeID
		productReq.SellerID = req.SellerID

		//se llama al service enviandole el producto a guardar
		productResp, err := p.productService.Save(productReq)

		//se responde error si hay
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewError(http.StatusInternalServerError, err.Error()))
			return
		}

		var resp response = response{}
		resp.Data = productResp
		//se responde un 201 con el producto ingresado
		ctx.JSON(201, resp.Data)

	}
}

//GetAll: devuelve una lista de productos
func (p *Product) GetAll() gin.HandlerFunc {

	type response struct {
		Data []domain.Product `json:"data"`
	}

	return func(ctx *gin.Context) {

		//se llama al service para traer una lista de productos
		products, err := p.productService.GetAll()

		//si hay error se devuelve
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewError(http.StatusInternalServerError, err.Error()))
			return
		}

		productosResp := response{}
		productosResp.Data = products

		//se devuelve un 200 con la lista de productos
		ctx.JSON(200, productosResp)
	}
}

//Get: devuelve un producto por id
func (p *Product) Get() gin.HandlerFunc {

	type response struct {
		Data domain.Product `json:"data"`
	}

	return func(ctx *gin.Context) {

		//se transforma el id ingresado por parametro a valor 'int'
		id, err := strconv.Atoi(ctx.Param("id"))

		//si hay un error en la conversion se devuelve
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewError(http.StatusBadRequest, err.Error()))
			return
		}

		//Se llama al service enviandole el id del producto para devolverlo
		producto, err := p.productService.Get(id)

		//Si hay error se devuelve
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewError(http.StatusNotFound, err.Error()))
			return
		}

		productosResp := response{}
		productosResp.Data = producto

		//se devuelve un 200 y el producto buscado
		ctx.JSON(200, productosResp)
	}
}

//Update: actualiza un producto completo
func (p *Product) Update() gin.HandlerFunc {

	type request struct {
		Description    string  `json:"description" binding:"required"`
		ExpirationRate int     `json:"expiration_rate" binding:"required"`
		FreezingRate   int     `json:"freezing_rate" binding:"required"`
		Height         float64 `json:"height" binding:"required"`
		Length         float64 `json:"length" binding:"required"`
		Netweight      float64 `json:"netweight" binding:"required"`
		ProductCode    string  `json:"product_code" binding:"required"`
		RecomFreezTemp float64 `json:"recommended_freezing_temperature" binding:"required"`
		Width          float64 `json:"width"  binding:"required"`
		ProductTypeID  int     `json:"product_type_id" binding:"required" validate:"number"`
		SellerID       int     `json:"seller_id" binding:"required" validate:"number"`
	}

	type response struct {
		Data domain.Product `json:"data"`
	}

	return func(ctx *gin.Context) {
		//se transforma el id ingresado por parametro a valor 'int'
		id, err := strconv.Atoi(ctx.Param("id"))

		//si hay un error en la conversion se devuelve
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewError(http.StatusBadRequest, err.Error()))
			return
		}

		var req request
		//si el JSON no contiene los campos necesarios se devuelve un 422
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		var productReq domain.Product
		//se asigna el id ingresado por parametro al producto
		productReq.ID = id
		productReq.Description = req.Description
		productReq.ExpirationRate = req.ExpirationRate
		productReq.FreezingRate = req.FreezingRate
		productReq.Height = req.Height
		productReq.Length = req.Length
		productReq.Netweight = req.Netweight
		productReq.ProductCode = req.ProductCode
		productReq.RecomFreezTemp = req.RecomFreezTemp
		productReq.Width = req.Width
		productReq.ProductTypeID = req.ProductTypeID
		productReq.SellerID = req.SellerID

		//se llama al service enviandole el producto actualizado
		product, err := p.productService.Update(productReq)

		//Si hay error se devuelve
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewError(http.StatusBadRequest, err.Error()))
			return
		}

		var resp response = response{}
		resp.Data = product
		//se devuelve un 200 y el producto actualizado
		ctx.JSON(200, resp.Data)
	}
}

//Delete: elimina un producto por id
func (p *Product) Delete() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		//se transforma el id ingresado por parametro a valor 'int'
		id, err := strconv.Atoi(ctx.Param("id"))

		//si hay un error en la conversion se devuelve
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewError(http.StatusBadRequest, err.Error()))
			return
		}
		//se llama al service enviandole el id del producto
		err = p.productService.Delete(id)

		//si hay error se devuelve
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewError(http.StatusNotFound, err.Error()))
			return
		}

		//se devuelve un 204
		ctx.JSON(204, "")
	}
}
