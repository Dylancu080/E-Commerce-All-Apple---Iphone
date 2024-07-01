package repository

import (
	"database/sql"
	"ecommerce/internal/models"
)

type CategoriaRepository struct {
	db *sql.DB
}

func NewCategoriaRepository(db *sql.DB) *CategoriaRepository {
	return &CategoriaRepository{db: db}
}

func (repo *CategoriaRepository) CreateCategoria(categoria *models.Categoria) error {
	query := "INSERT INTO categorias (nombre, descripcion) VALUES (?, ?)"
	_, err := repo.db.Exec(query, categoria.NombreCategoria, categoria.Descripcion)
	return err
}
