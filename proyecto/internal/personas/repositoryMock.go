package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/internal/models"
)

type RepositorySQLMock interface {
	Store(persona models.Persona) (models.Persona, error)
	GetOne(id int) models.Persona
	Update(persona models.Persona) (models.Persona, error)
	GetAll() ([]models.Persona, error)
	Delete(id int) error
	GetFullData() ([]models.Persona, error)

	GetOneWithContext(ctx context.Context, id int) (models.Persona, error)
}

type repositorySQLMock struct {
	db *sql.DB
}

func NewRepositorySQLMock(db *sql.DB) RepositorySQLMock {
	return &repositorySQLMock{db}
}

func (r *repositorySQLMock) Store(persona models.Persona) (models.Persona, error) {

	stmt, err := r.db.Prepare("INSERT INTO personas(nombre, apellido, edad) VALUES( ?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(persona.Nombre, persona.Apellido, persona.Edad)
	if err != nil {
		return models.Persona{}, err
	}

	idCreado, _ := result.LastInsertId()
	persona.ID = int(idCreado)

	return persona, nil
}

func (r *repositorySQLMock) GetOne(id int) models.Persona {

	var personaLeida models.Persona
	rows, err := r.db.Query("SELECT id, nombre,apellido, edad FROM personas WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return personaLeida
	}

	for rows.Next() {
		err = rows.Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad)
		if err != nil {
			log.Fatal(err)
			return personaLeida
		}

	}
	return personaLeida
}
func (r *repositorySQLMock) GetAll() ([]models.Persona, error) {
	var misPersonas []models.Persona
	var personaLeida models.Persona
	rows, err := r.db.Query("SELECT id, nombre, apellido, edad FROM personas")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		misPersonas = append(misPersonas, personaLeida)
	}
	return misPersonas, nil
}

func (r *repositorySQLMock) Update(persona models.Persona) (models.Persona, error) {

	stmt, err := r.db.Prepare("UPDATE personas SET nombre = ?, apellido = ?, edad = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(persona.Nombre, persona.Apellido, persona.Edad, persona.ID)
	if err != nil {
		return models.Persona{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Persona{}, errors.New("No se encontro la persona")
	}

	return persona, nil
}

func (r *repositorySQLMock) Delete(id int) error {

	stmt, err := r.db.Prepare("DELETE FROM personas WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return errors.New("No se encontro la persona")
	}
	return nil
}

func (r *repositorySQLMock) GetFullData() ([]models.Persona, error) {
	var misPersonas []models.Persona

	var personaLeida models.Persona
	rows, err := r.db.Query("select p.id,p.nombre, p.apellido, p.edad, c.nombre, c.nombrepais from personas p inner join ciudad c on p.idciudad = c.id")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad, &personaLeida.Domicilio.Nombre, &personaLeida.Domicilio.NombrePais)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		misPersonas = append(misPersonas, personaLeida)
	}
	return misPersonas, nil
}

func (r *repositorySQLMock) GetOneWithContext(ctx context.Context, id int) (models.Persona, error) {

	var personaLeida models.Persona
	// rows, err := db.QueryContext(ctx, "select sleep(30) from dual")
	rows, err := r.db.QueryContext(ctx, "SELECT id, nombre,apellido, edad FROM personas WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return personaLeida, err
	}

	for rows.Next() {
		err = rows.Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad)
		if err != nil {
			log.Fatal(err)
			return personaLeida, err
		}

	}
	return personaLeida, nil
}




func Test_sqlRepository_GetOne_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// mock.
	// mock.ExpectExec("INSERT INTO personas").WillReturnResult(sqlmock.NewResult(1, 1))
	columns := []string{"id", "nombre", "apellido", "edad"}
	rows := sqlmock.NewRows(columns)
	userId := 1
	rows.AddRow(userId, "", "", 0)
	mock.ExpectQuery("SELECT id, nombre,apellido, edad FROM personas WHERE id = ?").WithArgs(userId).WillReturnRows(rows)
	repository := NewRepositorySQLMock(db)
	user := models.Persona{
		ID: userId,
	}
	getResult := repository.GetOne(userId)
	// assert.Nil(t, getResult)
	// newPerso, err := repository.Store(user)
	// assert.Nil(t, err)
	assert.Equal(t, user.ID, getResult.ID)
	// getResult = repository.GetOne(userId)
	// assert.NoError(t, err)
	// assert.NotNil(t, getResult)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_sqlRepository_Store_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("").WillReturnResult(sqlmock.NewResult(8, 1))
	//.WillReturnResult(sqlmock.NewResult(1, 1))
	repository := NewRepositorySQLMock(db)
	userId := 8
	user := models.Persona{
		ID: userId,
	}
	// getResult := repository.GetOne(userId)
	// assert.Nil(t, getResult)
	newPerso, _ := repository.Store(user)
	// assert.Nil(t, err)
	assert.Equal(t, user.ID, newPerso.ID)
	// getResult = repository.GetOne(userId)
	// assert.NoError(t, err)
	// assert.NotNil(t, getResult)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_sqlRepository_Store_Failed_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO")
	mock.ExpectExec("INSERT INTO").WillReturnError(errors.New("No se puede insertar la persona"))
	//.WillReturnResult(sqlmock.NewResult(1, 1))
	repository := NewRepositorySQLMock(db)
	userId := 8
	user := models.Persona{
		ID: userId,
	}
	// getResult := repository.GetOne(userId)
	// assert.Nil(t, getResult)
	_, err = repository.Store(user)
	// assert.Nil(t, err)
	assert.Equal(t, "No se puede insertar la persona", err.Error())
	// getResult = repository.GetOne(userId)
	// assert.NoError(t, err)
	// assert.NotNil(t, getResult)
	assert.NoError(t, mock.ExpectationsWereMet())
}




func TestGetOk(t *testing.T) {
	db, err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	repository := NewRepositorySQLMock(db)
	defer db.Close()
	personaNueva := models.Persona{
		Nombre:   "Matias",
		Apellido: "Perez",
		Edad:     27,
	}

	personaCreada, _ := repository.Store(personaNueva)
	personaNueva.ID = personaCreada.ID
	personaCreada = repository.GetOne(personaNueva.ID)
	assert.Equal(t, personaNueva, personaCreada)
	assert.Equal(t, personaCreada.ID, personaNueva.ID)
}