package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
	warehouse "github.com/extmatperez/meli_bootcamp10_sprints/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp10_sprints/pkg/web"
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	warehouseService warehouse.Service
}

func NewWarehouse(w warehouse.Service) *Warehouse {
	return &Warehouse{
		warehouseService: w,
	}
}

//Get obtiene un warehouse de acuerdo al ID pasado como parámetro en la URL
func (w *Warehouse) Get() gin.HandlerFunc {
	type response struct {
		Data domain.Warehouse `json:"data"`
	}

	return func(c *gin.Context) {

		//Separo el ID que llega como parámetro y lo paso a entero
		id := c.Param("id")
		depId, _ := strconv.Atoi(id)

		//Llamo a la función Get del servicio. En caso de haber un error, lo devuelvo.
		deposito, err := w.warehouseService.Get(c, depId)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewError(http.StatusNotFound, err.Error()))
			return
		}
		//Devuelvo una respuesta correcta, junto al depósito solicitado.
		c.JSON(http.StatusOK, response{Data: deposito})

	}
}

//GetAll obtiene todos los warehouse existentes en la BD
func (w *Warehouse) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Warehouse `json:"data"`
	}

	return func(c *gin.Context) {

		//Creo slice de tipo Deposito (ID, Address, Telephone...)
		var depositos []domain.Warehouse

		//Llamo al servicio con el método. Chequeo si me llega un error.
		depositos, err := w.warehouseService.GetAll(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewError(http.StatusInternalServerError, err.Error()))
			return
		}

		//Si el largo del Slice es 0, no hay depositos cargados. Devuelvo el error correspondiente
		if len(depositos) == 0 {
			c.JSON(http.StatusInternalServerError, web.NewError(http.StatusInternalServerError, "no hay depositos almacenados"))
			return
		}
		c.JSON(http.StatusOK, response{Data: depositos})
	}
}

//Store da de alta un nuevo warehouse, de acuerdo a los datos enviados en el request
func (w *Warehouse) Store() gin.HandlerFunc {
	type request struct {
		Address            string `json:"address" binding:"required"`
		Telephone          string `json:"telephone" binding:"required"`
		WarehouseCode      string `json:"warehouse_code" binding:"required"`
		MinimunCapacity    int    `json:"minimun_capacity" binding:"required"`
		MinimunTemperature int    `json:"minimun_temperature" binding:"required"`
		SectionNumber      int    `json:"section_number" binding:"required"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		var req request

		//Si hay error en los campos requeridos, devuelvo código 422 con el error.
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, web.NewError(http.StatusUnprocessableEntity, err.Error()))
			return
		}

		//Ver si el codigo de deposito existe. Si existe mando un 404
		existe := w.warehouseService.Exists(c, req.WarehouseCode)
		if existe {
			c.JSON(http.StatusNotFound, web.NewError(http.StatusNotFound, "el código de depósito ya existe"))
			return
		}

		//Variable del struct
		var p domain.Warehouse

		//Datos struct
		p.Address = req.Address
		p.Telephone = req.Telephone
		p.WarehouseCode = req.WarehouseCode
		p.MinimunCapacity = req.MinimunCapacity
		p.MinimunTemperature = req.MinimunTemperature
		p.SectionNumber = req.SectionNumber

		//Llamo a la funcion para guardar el warehouse
		alta, err := w.warehouseService.Save(c, p)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewError(http.StatusBadRequest, err.Error()))
			return
		}

		//Voy a buscar lo creado y lo devuelvo en el response
		altaObj, err := w.warehouseService.Get(c, alta)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewError(http.StatusBadRequest, err.Error()))
			return
		}
		//Devuelvo 201 junto con el objeto creado
		c.JSON(http.StatusCreated, response{Data: altaObj})

	}
}

//Update es la función que nos permite actualizar los datos de un warehouse
func (w *Warehouse) Update() gin.HandlerFunc {
	type request struct {
		Address            string `json:"address" binding:"required"`
		Telephone          string `json:"telephone" binding:"required"`
		WarehouseCode      string `json:"warehouse_code" binding:"required"`
		MinimunCapacity    int    `json:"minimun_capacity" binding:"required"`
		MinimunTemperature int    `json:"minimun_temperature" binding:"required"`
		SectionNumber      int    `json:"section_number" binding:"required"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		var req request

		//Si hay error en los campos requeridos, devuelvo código 422 con el error. Todos son requeridos.
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(422, web.NewError(422, err.Error()))
			return
		}

		//Ver si el codigo de deposito existe
		existe := w.warehouseService.Exists(c, req.WarehouseCode)
		if !existe {
			c.JSON(http.StatusBadRequest, web.NewError(404, "el código de depósito no existe"))
			return
		}

		id := c.Param("id")
		depId, _ := strconv.Atoi(id)

		//Variable del struct
		var p domain.Warehouse

		//Datos struct
		p.ID = depId
		p.Address = req.Address
		p.Telephone = req.Telephone
		p.WarehouseCode = req.WarehouseCode
		p.MinimunCapacity = req.MinimunCapacity
		p.MinimunTemperature = req.MinimunTemperature
		p.SectionNumber = req.SectionNumber

		err := w.warehouseService.Update(c, p)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewError(http.StatusBadRequest, err.Error()))
			return
		}
		//Si la actualización fue exitosa, devuelvo
		c.JSON(200, response{Data: p})
	}
}

//Delete es una función que me permite eliminar un warehouse
func (w *Warehouse) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {

		//Paso el ID recibido como parámetro a INT y controlo que no haya errores en la conversion
		id := c.Param("id")
		wareId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewError(http.StatusNotFound, "invalid ID"))
			return
		}

		//Controlo si el warehouse existe. Para eso hago un Get, obtengo el objeto y luego controlo con EXIST(codigo de depósito como parámetro). 404 Si el cod dep no esiste.
		_, err = w.warehouseService.Get(c, wareId)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewError(http.StatusNotFound, err.Error()))
			return
		}
		// existe := w.warehouseService.Exists(c, obj.WarehouseCode)
		// if !existe {
		// 	c.JSON(http.StatusNotFound, web.NewError(http.StatusNotFound, "el código de depósito no existe"))
		// 	return
		// }

		//Si llegue hasta acá el elemento existe y puedo eliminar.
		err = w.warehouseService.Delete(c, wareId)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewError(http.StatusNotFound, err.Error()))
			return
		}
		//Ante una eliminación exitosa devuelvo Cod. 204
		c.JSON(http.StatusNoContent, nil)
	}
}
