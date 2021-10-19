package handler

import (
	"strconv"

	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/seller"
	"github.com/extmatperez/meli_bootcamp10_sprints/pkg/web"
	"github.com/gin-gonic/gin"
)

type Seller struct {
	sellerService seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{
		sellerService: s,
	}
}

// obtiene todos los sellers guardados en la base de datos
func (s *Seller) GetAll() gin.HandlerFunc {

	type response struct {
		Data []domain.Seller `json:"data"`
	}

	return func(c *gin.Context) {
		// trae de la base de datos los sellers
		sellers, err := s.sellerService.GetAll(c)
		if err != nil {
			e := web.NewError(500, err.Error())
			c.JSON(500, e)
			return
		}

		c.JSON(200, response{
			Data: sellers,
		})
	}
}

// obtiene seller por CID
func (s *Seller) Get() gin.HandlerFunc {

	type response struct {
		Data domain.Seller `json:"data"`
	}

	return func(c *gin.Context) {

		// valida si el cid es valido
		cid, err := strconv.Atoi(c.Param("cid"))
		if err != nil {
			c.JSON(400, web.NewError(400, "invalid cid. %v"))
			return
		}

		// trae de la base de datos el seller y en caso de que no exista devuelve un 404
		seller, err := s.sellerService.GetByCid(c, cid)
		if err != nil {
			c.JSON(404, web.NewErrorf(404, "seller with cid %d not found", cid))
			return
		}

		c.JSON(200, response{
			Data: seller,
		})
	}
}

// guarda un seller si el CID es unico
func (s *Seller) Store() gin.HandlerFunc {
	type response struct {
		Data domain.Seller `json:"data"`
	}

	return func(c *gin.Context) {
		// valida los campos de la request que tienen tag binding: "required" y tambien su zero value
		var req domain.SellerToSave
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(422, web.NewError(422, err.Error()))
			return
		}
		// valida que no exista el cid, si existe devuelve un error 409 conflict
		if exists := s.sellerService.ExistsCid(c, req.CID); exists {
			c.JSON(409, web.NewErrorf(409, "seller with cid %d already exists", req.CID))
			return
		}
		// guarda en la base de datos el seller
		storedSellerId, err := s.sellerService.Save(c, req)
		if err != nil {
			c.JSON(500, web.NewError(500, err.Error()))
			return
		}
		// trae de la base de datos el seller guardado
		storedSeller, err := s.sellerService.GetById(c, storedSellerId)
		if err != nil {
			c.JSON(500, web.NewError(500, err.Error()))
			return
		}

		c.JSON(201, response{
			Data: storedSeller,
		})
	}
}

// actualiza un seller si el CID es unico y existe
func (s *Seller) Update() gin.HandlerFunc {
	type response struct {
		Data domain.Seller `json:"data"`
	}

	return func(c *gin.Context) {
		// valida si el id es valido
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(422, web.NewError(422, "invalid id %v"))
			return
		}
		// valida los campos de la request que tienen tag binding: "required" y tambien su zero value
		var req domain.SellerToSave
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(422, web.NewError(422, err.Error()))
			return
		}
		// valida que el cid de la request no exista en la base de datos
		if exists := s.sellerService.ExistsCid(c, req.CID); exists {
			c.JSON(409, web.NewErrorf(409, "seller with cid %d already exists", req.CID))
			return
		}
		// valida que el id del seller exista
		if exists := s.sellerService.ExistsId(c, id); !exists {
			c.JSON(404, web.NewErrorf(404, "seller with id %d does not exists", id))
			return
		}
		// actualiza el seller en la base de datos
		if err = s.sellerService.Update(c, domain.Seller{
			ID:           id,
			SellerToSave: req,
		}); err != nil {
			c.JSON(500, web.NewError(500, err.Error()))
			return
		}
		// trae el seller actualizado de la base de datos
		storedSeller, err := s.sellerService.GetById(c, id)
		if err != nil {
			c.JSON(500, web.NewError(500, err.Error()))
			return
		}

		c.JSON(200, response{
			Data: storedSeller,
		})
	}
}

// elimina un seller por id si existe el seller en la base de datos
func (s *Seller) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {
		// valida si el id es valido
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, web.NewError(400, "invalid id %v"))
			return
		}
		// valida que el seller exista con ese id
		if exists := s.sellerService.ExistsId(c, id); !exists {
			c.JSON(404, web.NewErrorf(404, "seller with id %d does not exists", id))
			return
		}
		// elimina en la base de datos el seller
		if err = s.sellerService.Delete(c, id); err != nil {
			c.JSON(500, web.NewError(404, err.Error()))
			return
		}

		c.JSON(204, gin.H{})
	}
}
