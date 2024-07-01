//@autor: Dylan Curay
//@Fecha: 06 / 10 / 2024
//@Version: 1.0
//@Funcion: Definimos funciones de categorias para los productos

package models

import (
	"errors"
	"fmt"
)

// Definición de la estructura Categoria
type Categoria struct {
	CategoriaID     int    `json:"categoria_id"`
	NombreCategoria string `json:"nombre"`
	Descripcion     string `json:"descripcion"`
}

// Interfaz para categoría (opcional según la necesidad)
type CategoriaInterface interface {
	AgregarCategoria(categorias []Categoria, nuevaCategoria Categoria) []Categoria
	RemoverCategoria(categorias []Categoria, categoriaID int) ([]Categoria, error)
	ActualizarCategoria(categorias []Categoria, categoriaActualizada Categoria) ([]Categoria, error)
	LeerCategorias(categorias []Categoria)
}

// Función para agregar una nueva categoría
func (c *Categoria) AgregarCategoria(categorias []Categoria, nuevaCategoria Categoria) []Categoria {
	return append(categorias, nuevaCategoria)
}

// Función para remover una categoría por ID
func (c *Categoria) RemoverCategoria(categorias []Categoria, categoriaID int) ([]Categoria, error) {
	for i, categoria := range categorias {
		if categoria.CategoriaID == categoriaID {
			return append(categorias[:i], categorias[i+1:]...), nil
		}
	}
	return categorias, errors.New("Categoría para eliminar no encontrada")
}

// Función para actualizar una categoría por ID
func (c *Categoria) ActualizarCategoria(categorias []Categoria, categoriaActualizada Categoria) ([]Categoria, error) {
	for i, categoria := range categorias {
		if categoria.CategoriaID == categoriaActualizada.CategoriaID {
			categorias[i] = categoriaActualizada
			return categorias, nil
		}
	}
	return categorias, errors.New("Categoría para actualizar no encontrada")
}

// Método para imprimir las categorías (ejemplo básico)
func (c *Categoria) LeerCategorias(categorias []Categoria) {
	for _, cat := range categorias {
		fmt.Printf("ID: %d, Nombre: %s, Descripción: %s\n", cat.CategoriaID, cat.NombreCategoria, cat.Descripcion)
	}
}
