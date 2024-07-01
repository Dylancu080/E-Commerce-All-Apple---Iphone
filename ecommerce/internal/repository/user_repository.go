package repository

import (
	"database/sql"
	"ecommerce/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user *models.Usuarios) error {
	query := "INSERT INTO users (primer_nombre, segundo_nombre, primer_apellido, segundo_apellido, email, password, numero, genero, edad, fecha_nacimiento, es_admin) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := repo.db.Exec(query, user.PrimerNombre, user.SegundoNombre, user.PrimerApellido, user.SegundoApellido, user.Email, user.Password, user.Numero, user.Genero, user.Edad, user.FechaNacimiento, user.EsAdmin)
	return err
}
