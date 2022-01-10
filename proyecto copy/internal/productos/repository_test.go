package internal

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/w2GoPrueba/GoStorage/Ejemplo/proyecto/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestGetOne(t *testing.T) {
	//Arrange
	db := db.StorageDB
	repo := NewRepositoryProducto(db)

	productoCargada := repo.GetOne(1)
	fmt.Println(productoCargada)
	// assert.Equal(t, personaNueva.Nombre, personaCargada.Nombre)
	// assert.Equal(t, personaNueva.Apellido, personaCargada.Apellido)
	// assert.Nil(t, misPersonas)
}

func TestSumasPorCategorias(t *testing.T) {
	//Arrange
	db := db.StorageDB
	repo := NewRepositoryProducto(db)

	sumaPorCategorias, err := repo.SumaPorCategoria()
	assert.Nil(t, err)
	fmt.Printf("\n%v", sumaPorCategorias)
	fmt.Printf("\n%+v", sumaPorCategorias)
	fmt.Printf("\n%#v", sumaPorCategorias)
	// assert.Equal(t, personaNueva.Nombre, personaCargada.Nombre)
	// assert.Equal(t, personaNueva.Apellido, personaCargada.Apellido)
	// assert.Nil(t, misPersonas)
}

func TestSumasPorCategoriasTxDB(t *testing.T) {
	//Arrange
	db, err := db.InitDb()
	assert.Nil(t, err)
	repo := NewRepositoryProducto(db)

	sumaPorCategorias, err := repo.SumaPorCategoria()
	assert.Nil(t, err)
	fmt.Printf("\n%v", sumaPorCategorias)
	fmt.Printf("\n%+v", sumaPorCategorias)
	fmt.Printf("\n%#v", sumaPorCategorias)
	// assert.Equal(t, personaNueva.Nombre, personaCargada.Nombre)
	// assert.Equal(t, personaNueva.Apellido, personaCargada.Apellido)
	// assert.Nil(t, misPersonas)
}

func TestGetOneMock(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	rowsProductos := sqlmock.NewRows([]string{"id", "nombre", "precio", "idcategoria"})
	rowsProductos.AddRow(1, "Queso", 850, 2)

	rowsCategoria := sqlmock.NewRows([]string{"id", "nombre"})
	rowsCategoria.AddRow(1, "Fresquito")
	mock.ExpectQuery("SELECT id, nombre, precio, idcategoria FROM producto WHERE id = ?").WithArgs(1).WillReturnRows(rowsProductos)
	mock.ExpectQuery("SELECT id, nombre FROM categoria WHERE id = ?").WithArgs(2).WillReturnRows(rowsCategoria)

	repo := NewRepositoryProducto(db)

	productoCargada := repo.GetOne(1)

	fmt.Println(productoCargada)

	assert.NoError(t, mock.ExpectationsWereMet())
	// assert.Nil(t, misPersonas)
}
