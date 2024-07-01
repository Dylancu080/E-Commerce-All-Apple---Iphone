package repository

import (
	"database/sql"
	"ecommerce/internal/models"
)

type PedidoRepository struct {
	db *sql.DB
}

func NewPedidoRepository(db *sql.DB) *PedidoRepository {
	return &PedidoRepository{db: db}
}

func (repo *PedidoRepository) CreatePedido(pedido *models.Pedido) error {
	query := "INSERT INTO pedidos (usuario_id, fecha_pedido) VALUES (?, ?)"
	_, err := repo.db.Exec(query, pedido.UserID, pedido.FechaPedido)
	return err
}
