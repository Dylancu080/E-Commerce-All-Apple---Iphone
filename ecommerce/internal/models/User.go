//@autor: Dylan Curay
//@Fecha: 06 / 10 / 2024
//@Version: 1.0
//@Funcion: El usuario sera la cabeza para poder realizar todas las tareas, cada cuenta es distinta

package models

import (
	"errors"
	"fmt"
)

const claveMaestra = "claveMaestra123" // Clave maestra conocida por los programadores

// Estructura de Usuarios
type Usuarios struct {
	UserID          int
	PrimerNombre    string
	SegundoNombre   string
	PrimerApellido  string
	SegundoApellido string
	Email           string
	Password        string
	Numero          string
	Genero          string
	Edad            int
	FechaNacimiento string
	EsAdmin         bool
}

// Método para presentar los detalles del usuario
func (u *Usuarios) PresentarUsuario() {
	fmt.Printf("ID: %v\nPrimer Nombre: %v\nPrimer Apellido: %v\nSegundo Apellido: %v\nEmail: %v\nNúmero: %v\nGénero: %v\nEdad: %v\nFecha de Nacimiento: %v\nEs Admin: %t\n",
		u.UserID, u.PrimerNombre, u.PrimerApellido, u.SegundoApellido, u.Email,
		u.Numero, u.Genero, u.Edad, u.FechaNacimiento, u.EsAdmin)
}

// Función para crear un nuevo usuario
func NuevoUsuario(id int, primerNombre, segundoNombre, primerApellido, segundoApellido, email, password, numero, genero string, edad int, fechaNacimiento string, esAdmin bool, claveIngresada string) (Usuarios, error) {
	if esAdmin && claveIngresada != claveMaestra {
		return Usuarios{}, errors.New("clave maestra incorrecta, no se puede crear un administrador")
	}
	if primerNombre == "" || segundoNombre == "" || email == "" || password == "" || numero == "" || genero == "" || fechaNacimiento == "" {
		return Usuarios{}, errors.New("todos los campos deben ser llenados por favor")
	}

	usuario := Usuarios{
		UserID:          id,
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
	return usuario, nil
}

// Método para verificar si el usuario es admin
func (u Usuarios) EsAdministrador() bool {
	return u.EsAdmin
}

// Método para permitir que solo los administradores puedan agregar, actualizar o eliminar productos y categorías
func (u Usuarios) PermisoAdmin() error {
	if !u.EsAdmin {
		return errors.New("permiso denegado: no tienes acceso de administrador")
	}
	return nil
}

// Método para autenticar al usuario con una clave
func (u Usuarios) Autenticar(clave string) error {
	if u.Password != clave {
		return errors.New("clave incorrecta")
	}
	return nil
}
