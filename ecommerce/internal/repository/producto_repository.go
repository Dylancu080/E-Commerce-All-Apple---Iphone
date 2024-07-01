package repository

import (
	"database/sql"
	"ecommerce/internal/models"
)

type ProductoRepository struct {
	db *sql.DB
}

func NewProductoRepository(db *sql.DB) *ProductoRepository {
	return &ProductoRepository{db: db}
}

func (repo *ProductoRepository) CreateProducto(producto *models.Producto) error {
	query := "INSERT INTO productos (nombre, descripcion, precio, stock) VALUES (?, ?, ?, ?)"
	_, err := repo.db.Exec(query, producto.NombreProducto, producto.Descripcion, producto.Precio, producto.Stok)
	return err
}
