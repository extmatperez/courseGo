package internal

import (
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TM/proyecto/internal/models"
	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TM/proyecto/pkg/db"
)

const (
	GetPersonaById    = "SELECT id,nombre,apellido,edad FROM personas WHERE id = ?"
	InsertPersona     = "INSERT INTO personas(nombre, apellido, edad) VALUES( ?, ?, ? )"
	UpdatePersonaById = "UPDATE personas SET nombre = ?, apellido = ?, edad = ? WHERE id = ?"
)

type RepositorySQL interface {
	// GetAll() ([]models.Persona, error)
	Store(persona models.Persona) (models.Persona, error)
	// Update(id int, nombre string, apellido string, edad int) (models.Persona, error)
	// UpdateNombre(id int, nombre string) (models.Persona, error)
	// Delete(id int) error
	// //Store2(nuevaPersona Persona)(Persona, error)
	// LastId() (int, error)

	// Store(name, productType string, count int, price float64) (models.Product, error)
	GetOne(id int) models.Persona
	Update(persona models.Persona) (models.Persona, error)
	GetAll() ([]models.Persona, error)
	Delete(id int) error
}
type repositorySQL struct{}

func NewRepoSQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) Store(persona models.Persona) (models.Persona, error) {
	db := db.StorageDB                     // se inicializa la base
	stmt, err := db.Prepare(InsertPersona) // se prepara el SQL
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(persona.Nombre, persona.Apellido, persona.Edad) // retorna un sql.Result y un error
	if err != nil {
		return models.Persona{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Result devuelto en la ejecución obtenemos el Id insertado
	persona.ID = int(insertedId)

	return persona, nil
}
func (r *repositorySQL) GetOne(id int) models.Persona {
	var persona models.Persona
	db := db.StorageDB
	rows, err := db.Query(GetPersonaById, id)
	if err != nil {
		log.Println(err)
		return persona
	}
	for rows.Next() {
		if err := rows.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad); err != nil {
			log.Println(err.Error())
			return persona
		}
	}
	return persona
}

func (r *repositorySQL) Update(personaUpdated models.Persona) (models.Persona, error) {
	db := db.StorageDB                         // se inicializa la base
	stmt, err := db.Prepare(UpdatePersonaById) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	result, err := stmt.Exec(personaUpdated.Nombre, personaUpdated.Apellido, personaUpdated.Edad, personaUpdated.ID)
	if err != nil {
		return models.Persona{}, err
	}
	afectadas, _ := result.RowsAffected()
	if afectadas == 0 {
		return models.Persona{}, errors.New("persona no encontrada")
	}
	return personaUpdated, nil
}

func (r *repositorySQL) GetAll() ([]models.Persona, error) {
	var personas []models.Persona
	db := db.StorageDB
	rows, err := db.Query("SELECT id, nombre, apellido, edad FROM personas")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// se recorren todas las filas
	for rows.Next() {
		// por cada fila se obtiene un objeto del tipo Product
		var persona models.Persona
		if err := rows.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad); err != nil {
			log.Fatal(err)
			return nil, err
		}
		//se añade el objeto obtenido al slice products
		personas = append(personas, persona)
	}
	return personas, nil
}

func (r *repositorySQL) Delete(id int) error {
	db := db.StorageDB                                           // se inicializa la base
	stmt, err := db.Prepare("DELETE FROM personas WHERE id = ?") // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria

	result, err := stmt.Exec(id) // retorna un sql.Result y un error

	afectadas, _ := result.RowsAffected()
	if afectadas == 0 {
		return errors.New("persona no encontrada")
	}

	if err != nil {
		return err
	}

	return nil
}
