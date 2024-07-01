package service

import (
	"ecommerce/internal/models"
	"ecommerce/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(primerNombre, segundoNombre, primerApellido, segundoApellido, email, password, numero, genero string, edad int, fechaNacimiento string, esAdmin bool, claveIngresada string) (*models.Usuarios, error) {
	user := &models.Usuarios{
		PrimerNombre:    primerNombre,
		SegundoNombre:   segundoNombre,
		PrimerApellido:  primerApellido,
		SegundoApellido: segundoApellido,
		Email:           email,
		Password:        password,
		Numero:          numero,
		Genero:          genero,
		Edad:            edad,
		FechaNacimiento: fechaNacimiento,
		EsAdmin:         esAdmin,
	}
	err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
