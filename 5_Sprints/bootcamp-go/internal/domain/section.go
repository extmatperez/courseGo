package domain

// Section represents a structure where the warehouses are going to be located
type Section struct {
	ID                 int `json:"id"`
	SectionNumber      int `json:"section_number" binding:"required"`
	CurrentTemperature int `json:"current_temperature" binding:"required"`
	MinimumTemperature int `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int `json:"current_capacity" binding:"required"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int `json:"maximum_capacity" binding:"required"`
	WarehouseID        int `json:"warehouse_id" binding:"required"`
	ProductTypeID      int `json:"product_type_id" binding:"required"`
}
