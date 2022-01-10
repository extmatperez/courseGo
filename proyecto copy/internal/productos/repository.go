package internal

import (
	"database/sql"
	"log"

	internal "github.com/extmatperez/w2GoPrueba/GoStorage/Ejemplo/proyecto/internal/categoria"
	"github.com/extmatperez/w2GoPrueba/GoStorage/Ejemplo/proyecto/internal/models"
)

type RepositoryProducto interface {
	GetOne(id int) models.Producto
	SumaPorCategoria() ([]models.SumaByCategoria, error)
}

type repositoryProducto struct {
	db *sql.DB
}

func NewRepositoryProducto(db *sql.DB) RepositoryProducto {
	return &repositoryProducto{db}
}

func (r *repositoryProducto) GetOne(id int) models.Producto {

	var productoLeido models.Producto
	rows, err := r.db.Query("SELECT id, nombre, precio, idcategoria FROM producto WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return productoLeido
	}

	for rows.Next() {
		err = rows.Scan(&productoLeido.ID, &productoLeido.Nombre, &productoLeido.Precio, &productoLeido.Categoria.ID)
		if err != nil {
			log.Fatal(err)
			return productoLeido
		}
	}
	repo := internal.NewRepositoryCategoria(r.db)
	productoLeido.Categoria = repo.GetOne(productoLeido.Categoria.ID)
	return productoLeido
}

func (r *repositoryProducto) SumaPorCategoria() ([]models.SumaByCategoria, error) {

	var sumasPorCategorias []models.SumaByCategoria
	rows, err := r.db.Query("SELECT  c.nombre, SUM(p.precio) FROM producto p INNER JOIN categoria c on p.idcategoria = c.id GROUP BY c.nombre")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var sumaPorCategoria models.SumaByCategoria
		err = rows.Scan(&sumaPorCategoria.Nombre, &sumaPorCategoria.Suma)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		sumasPorCategorias = append(sumasPorCategorias, sumaPorCategoria)
	}

	return sumasPorCategorias, nil
}
