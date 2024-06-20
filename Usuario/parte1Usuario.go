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
	fmt.Printf("Id: %v\nId Usuario: %v\nPrimer Nombre: %v\nPrimerApellido: %v\nSegundoApellido: %v\nEmail: %v\n
					Contraseña: %v\nnumero: %v\nGenero: %v\nEdad: %v\nFechaNacimiento: %v\n ",
					 u.userID, u.primerNombre, u.segundoNombre, u.primerapellido, u.segundoapellido, u.email,
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
		
