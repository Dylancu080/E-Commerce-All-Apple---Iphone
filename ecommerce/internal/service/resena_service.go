package service

import (
	"ecommerce/internal/models"
	"ecommerce/internal/repository"
)

type ResenaService struct {
	resenaRepository *repository.ResenaRepository
}

func NewResenaService(resenaRepository *repository.ResenaRepository) *ResenaService {
	return &ResenaService{resenaRepository: resenaRepository}
}

func (s *ResenaService) CreateResena(productoID, calificacion int, comentario string) (*models.Resenas, error) {
	resena := &models.Resenas{
		ProductoID:   productoID,
		Calificacion: calificacion,
		Comentario:   comentario,
	}
	err := s.resenaRepository.CreateResena(resena)
	if err != nil {
		return nil, err
	}
	return resena, nil
}
