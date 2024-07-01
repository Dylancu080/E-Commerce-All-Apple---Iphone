package repository

import (
	"database/sql"
	"ecommerce/internal/models"
)

type ResenaRepository struct {
	db *sql.DB
}

func NewResenaRepository(db *sql.DB) *ResenaRepository {
	return &ResenaRepository{db: db}
}

func (repo *ResenaRepository) CreateResena(resena *models.Resenas) error {
	query := "INSERT INTO resenas (producto_id, calificacion, comentario) VALUES (?, ?, ?)"
	_, err := repo.db.Exec(query, resena.ProductoID, resena.Calificacion, resena.Comentario)
	return err
}
