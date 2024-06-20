//@autor: Dylan Curay
//@Fecha: 06 / 10 / 2024
//@Version: 1.0
//@Funcion: Definimos funciones de categorias para los productos

package categoria

import (
	"errors"
	"fmt"
)

type Categoria struct {
	categoriaID     int
	nombreCategoria string
}

func agregarCategoria(categoria []Categoria, nuevaCategoria Categoria) []Categoria {
	return append(categoria, nuevaCategoria)

}

func removerCategoria(categoria []Categoria, categoriaID int) ([]Categoria, error) {
	for a, category := range categoria {
		if category.categoriaID == categoriaID {
			return append(categoria[:a], categoria[a+1:]...), nil
		}
	}
	return categoria, errors.New("Categoria para eliminar no encontrada")
}

func actualizatCategoria(categoria []Categoria, categoriaActualizada Categoria) ([]Categoria, error) {
	for a, category := range categoria {
		if category.categoriaID == categoriaActualizada.categoriaID {
			categoria[a] = categoriaActualizada
			return categoria, nil
		}
	}
	return categoria, errors.New("Categoria para actualizar no encontrada")
}

func (d Categoria) leerCategorias() {
	fmt.Printf("ID: %d, nombre: %s", d.categoriaID, d.nombreCategoria)
}
