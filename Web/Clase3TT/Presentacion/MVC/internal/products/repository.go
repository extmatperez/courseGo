package internal

type Product struct {
	ID      int     `json:"id"`
	Name    string  `json:"nombre"`
	Type    string  `json:"tipo"`
	Count   int     `json:"cantidad"`
	Price   float64 `json:"precio"`
	Estado  bool    `json:"estado"`
	Estado2 bool    `json:"estado2"`
}

var ps []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}
func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {
	p := Product{id, nombre, tipo, cantidad, precio, true, false}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}
