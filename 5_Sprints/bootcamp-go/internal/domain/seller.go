package domain

type SellerToSave struct {
	CID         int    `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityID  int    `json:"locality_id" binding:"required"`
}

type Seller struct {
	ID int `json:"id"`
	SellerToSave
}
