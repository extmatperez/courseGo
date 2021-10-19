package domain

type Employee struct {
	ID           int    `json:"id,omitempty"`
	CardNumberID string `json:"card_number_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	WarehouseID  int    `json:"warehouse_id,omitempty"`
}
