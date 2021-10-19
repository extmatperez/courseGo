package internal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/extmatperez/courseGo/Web/Clase3TT/Practica/pkg/store"
)

type Usuario struct {
	ID              int     `json:"id"`
	Nombre          string  `json:"nombre" binding:"required"`
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
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, FechaDeCreacion string) (Usuario, error)
	LastID() (int, error)
	UpdateNombre(id int, nombreNuevo string) (Usuario, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

// func (r *repository) GetAll() ([]Usuario, error) {
// 	return misUsuarios, nil
// }

func (r *repository) GetAll() ([]Usuario, error) {
	var ps []Usuario
	r.db.Read(&ps)
	return ps, nil
}
func (r *repository) LoadFile() ([]Usuario, error) {
	usr := "E:/Google Drive/DIGITAL HOUSE/Capacitacion GO/GitHub/courseGo/Web/Clase2TT/Practica/Ejercicio1/internal/usuarios/archivo/usuarios.json"
	data, err := os.ReadFile(usr)
	fmt.Println("Holis", len(data))
	if err != nil {
		return nil, err
	} else {
		var usuarios []Usuario
		json.Unmarshal(data, &usuarios)
		// ctxt.String(200, string(data))
		return usuarios, nil
	}
}

// func (r *repository) LastID() (int, error) {
// 	return lastID, nil
// }
func (r *repository) LastID() (int, error) {

	var ps []Usuario
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].ID, nil

}

// func (r *repository) Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (Usuario, error) {
// 	p := Usuario{id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion}
// 	misUsuarios = append(misUsuarios, p)
// 	lastID = p.ID
// 	return p, nil
// }

func (r *repository) Store(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (Usuario, error) {
	var ps []Usuario
	r.db.Read(&ps)
	p := Usuario{id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Usuario{}, err
	}
	return p, nil
}

func (r *repository) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (Usuario, error) {
	p := Usuario{id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion}
	updated := false

	for i := range misUsuarios {
		if misUsuarios[i].ID == id {
			p.ID = id
			misUsuarios[i] = p
			updated = true
		}
	}

	if !updated {
		return Usuario{}, fmt.Errorf("Usuario %d no encontrado...", id)
	}

	return p, nil
}
func (r *repository) UpdateNombre(id int, nombreNuevo string) (Usuario, error) {
	var user Usuario
	updated := false

	for i := range misUsuarios {
		if misUsuarios[i].ID == id {
			misUsuarios[i].Nombre = nombreNuevo
			updated = true
			user = misUsuarios[i]
		}
	}

	if !updated {
		return Usuario{}, fmt.Errorf("Usuario %d no encontrado...", id)
	}

	return user, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range misUsuarios {
		if misUsuarios[i].ID == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("Usuario %d no encontrado...", id)
	}
	misUsuarios = append(misUsuarios[:index], misUsuarios[index+1:]...)
	return nil
}
