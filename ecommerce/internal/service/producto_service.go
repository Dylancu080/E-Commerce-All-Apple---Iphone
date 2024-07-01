package service

import (
	"ecommerce/internal/models"
	"ecommerce/internal/repository"
)

type ProductoService struct {
	productoRepository *repository.ProductoRepository
}

func NewProductoService(productoRepository *repository.ProductoRepository) *ProductoService {
	return &ProductoService{productoRepository: productoRepository}
}

func (s *ProductoService) CreateProducto(productoID int, nombreProducto, descripcion string, precio float64, stock int) (*models.Producto, error) {
	producto := &models.Producto{
		ProductoID:     productoID,
		NombreProducto: nombreProducto,
		Descripcion:    descripcion,
		Precio:         precio,
		Stok:           stock,
	}
	err := s.productoRepository.CreateProducto(producto)
	if err != nil {
		return nil, err
	}
	return producto, nil
}
