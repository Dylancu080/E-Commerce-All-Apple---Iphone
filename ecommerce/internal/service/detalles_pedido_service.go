package service

import (
	"ecommerce/internal/models"
	"ecommerce/internal/repository"
)

type DetallesPedidoService struct {
	detallesPedidoRepository *repository.DetallesPedidoRepository
}

func NewDetallesPedidoService(detallesPedidoRepository *repository.DetallesPedidoRepository) *DetallesPedidoService {
	return &DetallesPedidoService{detallesPedidoRepository: detallesPedidoRepository}
}

func (s *DetallesPedidoService) CreateDetallesPedido(pedidoID, productoID, cantidad int) (*models.Detalles, error) {
	detallesPedido := &models.Detalles{
		PedidoID:   pedidoID,
		ProductoID: productoID,
		Cantidad:   cantidad,
	}
	err := s.detallesPedidoRepository.CreateDetallesPedido(detallesPedido)
	if err != nil {
		return nil, err
	}
	return detallesPedido, nil
}
