package service

import (
	"ecommerce/internal/models"
	"ecommerce/internal/repository"
)

type CategoriaService struct {
	categoriaRepository *repository.CategoriaRepository
}

func NewCategoriaService(categoriaRepository *repository.CategoriaRepository) *CategoriaService {
	return &CategoriaService{categoriaRepository: categoriaRepository}
}

func (s *CategoriaService) CreateCategoria(nombre, descripcion string) (*models.Categoria, error) {
	categoria := &models.Categoria{
		NombreCategoria: nombre,
		Descripcion:     descripcion,
	}
	err := s.categoriaRepository.CreateCategoria(categoria)
	if err != nil {
		return nil, err
	}
	return categoria, nil
}
