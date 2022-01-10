package internal

import (
	"database/sql"
	"log"

	"github.com/extmatperez/w2GoPrueba/GoStorage/Ejemplo/proyecto/internal/models"
)

type RepositoryCategoria interface {
	GetOne(id int) models.Categoria
}

type repositoryCategoria struct {
	db *sql.DB
}

func NewRepositoryCategoria(db *sql.DB) RepositoryCategoria {
	return &repositoryCategoria{db}
}

func (r *repositoryCategoria) GetOne(id int) models.Categoria {

	var categoriaLeida models.Categoria
	rows, err := r.db.Query("SELECT id, nombre FROM categoria WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return categoriaLeida
	}

	for rows.Next() {
		err = rows.Scan(&categoriaLeida.ID, &categoriaLeida.Nombre)
		if err != nil {
			log.Fatal(err)
			return categoriaLeida
		}

	}
	return categoriaLeida
}
