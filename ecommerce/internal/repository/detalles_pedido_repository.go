package repository

import (
	"database/sql"
	"ecommerce/internal/models"
)

type DetallesPedidoRepository struct {
	db *sql.DB
}

// Constructor para el repositorio de detalles de pedido
func NewDetallesPedidoRepository(db *sql.DB) *DetallesPedidoRepository {
	return &DetallesPedidoRepository{db: db}
}

// Método para crear un detalle de pedido en la base de datos
func (repo *DetallesPedidoRepository) CreateDetallesPedido(detallesPedido *models.Detalles) error {
	query := "INSERT INTO detalles_pedidos (pedido_id, producto_id, cantidad, precio_unitario) VALUES (?, ?, ?, ?)"
	_, err := repo.db.Exec(query, detallesPedido.PedidoID, detallesPedido.ProductoID, detallesPedido.Cantidad, detallesPedido.Precio)
	return err
}
