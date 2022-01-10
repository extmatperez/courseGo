package internal

import "github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TM/proyecto/internal/models"

type ServiceSQL interface {
	GetAll() ([]models.Persona, error)
	Store(nombre, apellido string, edad int) (models.Persona, error)
	GetOne(id int) models.Persona
	Update(personaUpdated models.Persona) (models.Persona, error)
	// UpdateNombre(id int, nombre string) (Persona, error)
	Delete(id int) error
	// Sum(prices ...float64) float64
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewserviceSQL(repository RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repository}
}

func (ser *serviceSQL) GetAll() ([]models.Persona, error) {
	personas, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return personas, nil
}

func (ser *serviceSQL) Store(nombre, apellido string, edad int) (models.Persona, error) {
	//ultimoId, err := ser.repository.LastId()

	// if err != nil {
	// 	return models.Persona{}, err
	// }

	newPersona := models.Persona{Nombre: nombre, Apellido: apellido, Edad: edad}
	per, err := ser.repository.Store(newPersona)

	if err != nil {
		return models.Persona{}, err
	}

	return per, nil
}

func (ser *serviceSQL) GetOne(id int) models.Persona {
	return ser.repository.GetOne(id)
}

func (ser *serviceSQL) Update(personaUpdated models.Persona) (models.Persona, error) {
	return ser.repository.Update(personaUpdated)
}

// func (ser *serviceSQL) UpdateNombre(id int, nombre string) (Persona, error) {
// 	return ser.repository.UpdateNombre(id, nombre)
// }

func (ser *serviceSQL) Delete(id int) error {
	return ser.repository.Delete(id)
}

// func (s *serviceSQL) Sum(prices ...float64) float64 {
// 	var price float64
// 	for _, p := range prices {
// 		price += p
// 	}
// 	return price
// }
