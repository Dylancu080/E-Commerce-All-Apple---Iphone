package service

import (
	"ecommerce/internal/models"
	"ecommerce/internal/repository"
)

type EntregaPedidoService struct {
	entregaPedidoRepository *repository.EntregaPedidoRepository
}

func NewEntregaPedidoService(entregaPedidoRepository *repository.EntregaPedidoRepository) *EntregaPedidoService {
	return &EntregaPedidoService{entregaPedidoRepository: entregaPedidoRepository}
}

func (s *EntregaPedidoService) CreateEntregaPedido(pedidoID int, deliveryAddress string) (*models.Entrega, error) {
	entregaPedido := &models.Entrega{
		PedidoID:        pedidoID,
		DeliveryAddress: deliveryAddress,
	}
	err := s.entregaPedidoRepository.CreateEntregaPedido(entregaPedido)
	if err != nil {
		return nil, err
	}
	return entregaPedido, nil
}
