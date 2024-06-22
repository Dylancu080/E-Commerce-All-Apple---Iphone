package resena

import (
	"time"
)

type Resena struct {
	ID           int
	ProductoID   int
	Calificacion int
	Comentario   string
	Fecha        time.Time
}

func agregar(resenas []Resena, nuevaResena Resena) []Resena {
	return append(resenas, nuevaResena)
}

func remover(resenas []Resena, id int) []Resena {
	for i, resena := range resenas {
		if resena.ID == id {
			return append(resenas[:i], resenas[i+1:]...)
		}
	}
	return resenas
}
