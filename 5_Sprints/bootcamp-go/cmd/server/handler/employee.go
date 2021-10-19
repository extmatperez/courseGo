package handler

import (
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/employee"
	"github.com/extmatperez/meli_bootcamp10_sprints/pkg/web"
	"github.com/gin-gonic/gin"
)

type Employee struct {
	employeeService employee.Service
}

//NewEmployee
//Create handler
//@return *Employee
func NewEmployee(e employee.Service) *Employee {
	return &Employee{
		employeeService: e,
	}
}

//Get
//Get one employee
//@return gin.HandlerFunc for Routing
func (e *Employee) Get() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	//get one employee
	return func(c *gin.Context) {
		card_number_id := c.Param("card_number_id")
		employee, err := e.employeeService.Get(card_number_id)
		if err != nil {
			c.JSON(400, web.NewError(400, err.Error()))
			return
		}

		c.JSON(200, response{employee})
	}
}

//GetAll
//Get all employees
//@return gin.HandlerFunc for Routing
func (e *Employee) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Employee `json:"data"`
	}
	//Get all employees
	return func(c *gin.Context) {
		employees, err := e.employeeService.GetAll()
		if err != nil {
			c.JSON(400, web.NewError(400, err.Error()))
			return
		}

		c.JSON(200, response{employees})

	}
}

//Store
//add employee
//@return gin.HandlerFunc for Routing
func (e *Employee) Store() gin.HandlerFunc {
	type request struct {
		CardNumberID string `json:"card_number_id" binding:"required"`
		FirstName    string `json:"first_name" binding:"required"`
		LastName     string `json:"last_name" binding:"required"`
		WarehouseID  int    `json:"warehouse_id" `
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	//add employee
	return func(c *gin.Context) {
		var req request

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(400, web.NewError(400, err.Error()))
			return
		}

		employee, err := e.employeeService.Store(req.CardNumberID, req.FirstName, req.LastName, req.WarehouseID)

		if err != nil {
			c.JSON(422, web.NewError(422, err.Error()))
			return
		}

		c.JSON(201, response{employee})

	}
}

//Update
//update employee
//@return gin.HandlerFunc for Routing
func (e *Employee) Update() gin.HandlerFunc {
	type request struct {
		FirstName   string `json:"first_name"  binding:"required"`
		LastName    string `json:"last_name" binding:"required"`
		WarehouseID int    `json:"warehouse_id" binding:"required"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}
	//update employee
	return func(c *gin.Context) {
		var req request

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(400, web.NewError(400, err.Error()))
			return
		}
		card_number_id := c.Param("card_number_id")

		employee, err := e.employeeService.Update(card_number_id, req.FirstName, req.LastName, req.WarehouseID)

		if err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		}

		c.JSON(200, response{employee})

	}
}

//Delete
//delete employee
//@return gin.HandlerFunc for Routing
func (e *Employee) Delete() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}
	//delete employee
	return func(c *gin.Context) {
		card_number_id := c.Param("card_number_id")

		err := e.employeeService.Delete(card_number_id)

		if err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		}

		c.JSON(204, response{"el empleado ha sido eliminado con exito"})

	}
}
