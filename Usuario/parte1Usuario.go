//@autor: Dylan Curay
//@Fecha: 06 / 10 / 2024
//@Version: 1.0
//@Funcion: El usuario sera la cabeza para poder realizar todas las tareas, cada cuenta es distinta

package usuario

import (
	"fmt"
)

type Usuarios struct {
	userID          int
	primerNombre    string
	segundoNombre   string
	primerApellido  string
	segundoApellido string
	email           string
	password        string
	numero          string
	genero          string
	edad            int
	fechaNacimiento string
}

func (u *Usuarios) PresentarUsuario() {
	fmt.Printf("Id: %v\nId Usuario: %v\nPrimer Nombre: %v\nPrimer Apellido: %v\nSegundo Apellido: %v\nEmail: %v\nContraseña: %v\nNúmero: %v\nGénero: %v\nEdad: %v\nFecha de Nacimiento: %v\n",
		u.userID, u.userID, u.primerNombre, u.primerApellido, u.segundoApellido, u.email,
		u.password, u.numero, u.genero, u.edad, u.fechaNacimiento)
}
