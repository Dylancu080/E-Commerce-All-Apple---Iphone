//@autor: Dylan Curay
//@Fecha: 06 / 10 / 2024
//@Version: 1.0
//@Funcion: El usuario sera la cabeza para poder realizar todas las tareas, cada cuenta es distinta

package usuario

import (
	"errors"
	"fmt"
	"main/Usuario"
)

type Usuarios struct {
	userID int
	primerNombre string
	segundoNombre string
	primerApellido string
	segundoApellido string
	email string
	password string
	numero string
	genero string
	edad int
	fechaNacimiento string
}
func (u *Usuarios) PresentarUsuario() {
    fmt.Printf("Id: %v\nId Usuario: %v\nPrimer Nombre: %v\nPrimer Apellido: %v\nSegundo Apellido: %v\nEmail: %v\nContraseña: %v\nNúmero: %v\nGénero: %v\nEdad: %v\nFecha de Nacimiento: %v\n",
        u.userID, u.userID, u.primerNombre, u.primerApellido, u.segundoApellido, u.email,
        u.password, u.numero, u.genero, u.edad, u.fechaNacimiento)
}


func NuevoUsuario(id int, primerNombre, segundoNombre, email, password, numero, genero string, 
	edad int, fechaNacimiento string) (Usuarios, error) {
		if primerNombre == "" || segundoNombre== "" || email== "" || password== "" || numero== "" || genero== "" || edad== "" || fechaNacimiento== "" {
			return Usuarios{}, errors.New("Todos los campos deben ser llenados por favor")
		}
		
		usuario := Usuarios{ 
			userID: id,
			primerNombre:  nombre,
			segundoNombre: nombre2,
			primerApellido: apellido
			segundoApellido: segundoApellido,
			email: email,
			password: clave,
			numero: numero,          
			genero: genero,
			edad: edad,
			fechaNacimiento: fechaNacimiento,
		
		}
		return usuario, nil

	}
		
