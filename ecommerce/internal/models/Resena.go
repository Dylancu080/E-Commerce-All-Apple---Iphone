//@autor: Dylan Curay
//@Fecha: 06 / 10 / 2024
//@Version: 1.0
//@Funcion: Damos las opciones para que el usuario deje una opinión sobre los productos que compro

package models

import (
	"errors"
	"fmt"
	"time"
)

type Resenas struct {
	ResenaID     int
	ProductoID   int
	Calificacion int
	Comentario   string
	FechaResena  time.Time
}

func AgregarResena(resenias []Resenas, nuevaResena Resenas) []Resenas {
	return append(resenias, nuevaResena)
}

func RemoverResena(resenias []Resenas, resenaID int) ([]Resenas, error) {
	for i, resena := range resenias {
		if resena.ResenaID == resenaID {
			return append(resenias[:i], resenias[i+1:]...), nil
		}
	}
	return resenias, errors.New("reseña no encontrada")
}

func (r Resenas) ImprimirResena() {
	fmt.Printf("Reseña ID: %d, Producto ID: %d, Calificación: %d, Comentario: %s, Fecha: %s\n",
		r.ResenaID, r.ProductoID, r.Calificacion, r.Comentario, r.FechaResena.Format("02-01-2006"))
}
