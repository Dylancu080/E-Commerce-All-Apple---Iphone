package repository

import (
	"database/sql"
	"ecommerce/internal/models"
)

type EntregaPedidoRepository struct {
	db *sql.DB
}

func NewEntregaPedidoRepository(db *sql.DB) *EntregaPedidoRepository {
	return &EntregaPedidoRepository{db: db}
}

func (repo *EntregaPedidoRepository) CreateEntregaPedido(entregaPedido *models.Entrega) error {
	query := "INSERT INTO entregas_pedidos (pedido_id, direccion_entrega) VALUES (?, ?)"
	_, err := repo.db.Exec(query, entregaPedido.PedidoID, entregaPedido.DeliveryAddress)
	return err
}
