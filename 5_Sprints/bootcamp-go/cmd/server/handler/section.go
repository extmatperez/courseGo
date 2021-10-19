package handler

import (
	"strconv"

	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/internal/section"
	"github.com/extmatperez/meli_bootcamp10_sprints/pkg/web"

	"github.com/gin-gonic/gin"
)

// Structure representing a Section
type Section struct {
	sectionService section.Service
}

// This creates a new Section handler
// Params: section.Service to get the service associated to the handler
// Returns a new Section handler
func NewSection(s *section.Service) *Section {
	return &Section{
		sectionService: *s,
	}
}

// Gets all the section stored in the database
// Returns all sections in a JSON response
func (s *Section) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Section `json:"data"`
	}
	return func(ctx *gin.Context) {
		sections, err := s.sectionService.GetAll(ctx)
		if err != nil {
			ctx.JSON(404, web.NewError(404, err.Error()))
		}
		ctx.JSON(200, response{
			Data: sections,
		})
	}
}

// Gets a specific section based on the section_number
// Returns specific section with the id passed as path params
func (s *Section) Get() gin.HandlerFunc {
	type response struct {
		Data domain.Section `json:"data"`
	}
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		secId, idErr := strconv.Atoi(id)
		if idErr != nil {
			ctx.JSON(400, web.NewError(400, idErr.Error()))
			return
		}

		section, err := s.sectionService.Get(ctx, secId)
		if err != nil {
			ctx.JSON(404, web.NewError(404, err.Error()))
			return
		}
		ctx.JSON(200, response{
			Data: section,
		})
	}
}

// Inserts a new section into the database with all of its fields
// Section has fields such as: SectionNumber, CurrentTemperature, MinTemperature, CurrentCapacity, MinCapacity, MaxCapacity, WarehouseID, ProductTypeID
func (s *Section) Store() gin.HandlerFunc {
	type request struct {
		SectionNumber      int `json:"section_number"`
		CurrentTemperature int `json:"current_temperature"`
		MinTemperature     int `json:"minimum_temperature"`
		CurrentCapacity    int `json:"current_capacity"`
		MinCapacity        int `json:"minimum_capacity"`
		MaxCapacity        int `json:"maximum_capacity"`
		WarehouseID        int `json:"warehouse_id"`
		ProductTypeID      int `json:"product_type_id"`
	}

	type response struct {
		Data domain.Section `json:"data"`
	}

	return func(ctx *gin.Context) {
		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(422, web.NewError(422, err.Error()))
			return
		}
		if req.SectionNumber == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.CurrentTemperature == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.MinTemperature == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.CurrentCapacity == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.MinCapacity == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.MaxCapacity == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.WarehouseID == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.ProductTypeID == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}

		newSection := domain.Section{
			SectionNumber:      req.SectionNumber,
			CurrentTemperature: req.CurrentTemperature,
			MinimumTemperature: req.MinTemperature,
			CurrentCapacity:    req.CurrentCapacity,
			MinimumCapacity:    req.MinCapacity,
			MaximumCapacity:    req.MaxCapacity,
			WarehouseID:        req.WarehouseID,
			ProductTypeID:      req.ProductTypeID,
		}
		secId, err := s.sectionService.Save(ctx, newSection)
		if err != nil {
			ctx.JSON(409, web.NewError(409, err.Error()))
			return
		}

		newSection.ID = secId

		ctx.JSON(201, response{
			Data: newSection,
		})
	}
}

// Updates all of the fields that a Section passed by path params has
func (s *Section) Update() gin.HandlerFunc {

	type request struct {
		SectionNumber      int `json:"section_number"`
		CurrentTemperature int `json:"current_temperature"`
		MinTemperature     int `json:"minimum_temperature"`
		CurrentCapacity    int `json:"current_capacity"`
		MinCapacity        int `json:"minimum_capacity"`
		MaxCapacity        int `json:"maximum_capacity"`
		WarehouseID        int `json:"warehouse_id"`
		ProductTypeID      int `json:"product_type_id"`
	}

	type response struct {
		SectionNumber      int `json:"section_number"`
		CurrentTemperature int `json:"current_temperature"`
		MinTemperature     int `json:"minimum_temperature"`
		CurrentCapacity    int `json:"current_capacity"`
		MinCapacity        int `json:"minimum_capacity"`
		MaxCapacity        int `json:"maximum_capacity"`
		WarehouseID        int `json:"warehouse_id"`
		ProductTypeID      int `json:"product_type_id"`
	}

	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		secId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, web.NewError(400, err.Error()))
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(422, web.NewError(422, err.Error()))
			return
		}
		if req.SectionNumber == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.CurrentTemperature == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.MinTemperature == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.CurrentCapacity == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.MinCapacity == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.MaxCapacity == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.WarehouseID == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}
		if req.ProductTypeID == 0 {
			ctx.JSON(422, web.NewError(422, "not all the fields were sent"))
			return
		}

		newSection := domain.Section{
			SectionNumber:      req.SectionNumber,
			CurrentTemperature: req.CurrentTemperature,
			MinimumTemperature: req.MinTemperature,
			CurrentCapacity:    req.CurrentCapacity,
			MinimumCapacity:    req.MinCapacity,
			MaximumCapacity:    req.MaxCapacity,
			WarehouseID:        req.WarehouseID,
			ProductTypeID:      req.ProductTypeID,
		}
		err = s.sectionService.Update(ctx, newSection, secId)
		if err != nil {
			ctx.JSON(404, web.NewError(404, err.Error()))
			return
		}

		ctx.JSON(200, response{
			newSection.SectionNumber,
			newSection.CurrentTemperature,
			newSection.MinimumTemperature,
			newSection.CurrentCapacity,
			newSection.MinimumCapacity,
			newSection.MaximumCapacity,
			newSection.WarehouseID,
			newSection.ProductTypeID,
		})
	}
}

// Deletes a section held within the database
func (s *Section) Delete() gin.HandlerFunc {
	type response struct {
		Msg string `json:"msg"`
	}
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		secId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, web.NewError(400, err.Error()))
			return
		}
		err = s.sectionService.Delete(ctx, secId)
		if err != nil {
			ctx.JSON(404, web.NewError(404, err.Error()))
			return
		}
		ctx.JSON(204, response{Msg: "Section deleted correctly"})
	}
}
