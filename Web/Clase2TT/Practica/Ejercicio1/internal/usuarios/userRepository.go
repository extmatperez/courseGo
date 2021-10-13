package internal

type Usuario struct {
	ID              int     `json:"id"`
	Nombre          string  `json:"nombre"`
	Apellido        string  `json:"apellido"`
	Email           string  `json:"email"`
	Edad            int     `json:"edad"`
	Altura          float64 `json:"altura"`
	Activo          bool    `json:"activo"`
	FechaDeCreacion string  `json:"fechaDeCreacion"`
}

var misUsuarios []Usuario
var lastID int

type Repository interface {
	GetAll() ([]Usuario, error)
	LoadFile() ([]Usuario, error)
	Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, FechaDeCreacion string) (Usuario, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}
func (r *repository) GetAll() ([]Usuario, error) {
	return misUsuarios, nil
}
func (r *repository) LoadFile() ([]Usuario, error) {

	return misUsuarios, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (Usuario, error) {
	p := Usuario{id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion}
	misUsuarios = append(misUsuarios, p)
	lastID = p.ID
	return p, nil
}
