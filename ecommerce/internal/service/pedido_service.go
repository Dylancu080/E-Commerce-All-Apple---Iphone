package service

import (
	"ecommerce/internal/models"
	"ecommerce/internal/repository"
)

type PedidoService struct {
	pedidoRepository *repository.PedidoRepository
}

func NewPedidoService(pedidoRepository *repository.PedidoRepository) *PedidoService {
	return &PedidoService{pedidoRepository: pedidoRepository}
}

func (s *PedidoService) CreatePedido(usuarioID int) (*models.Pedido, error) {
	pedido := &models.Pedido{
		UserID: usuarioID,
	}
	err := s.pedidoRepository.CreatePedido(pedido)
	if err != nil {
		return nil, err
	}
	return pedido, nil
}
