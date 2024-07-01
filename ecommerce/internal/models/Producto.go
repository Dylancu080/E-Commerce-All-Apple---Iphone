package models

import (
	"errors"
	"fmt"
)

type Producto struct {
	ProductoID     int
	UserID         int
	NombreProducto string
	Descripcion    string
	Precio         float64
	Categoria      string
	Stok           int
}

type ProductoInterface interface {
	AgregarProducto(productos []Producto, nuevoProducto Producto) []Producto
	ActualizarProducto(productos []Producto, productoActualizado Producto) ([]Producto, error)
	RemoverProducto(productos []Producto, productoID int) ([]Producto, error)
	BuscarProductoNombre(productos []Producto, nombreProducto string) (Producto, error)
}

func AgregarProducto(productos []Producto, nuevoProducto Producto) []Producto {
	return append(productos, nuevoProducto)
}

func ActualizarProducto(productos []Producto, productoActualizado Producto) ([]Producto, error) {
	for i, producto := range productos {
		if producto.ProductoID == productoActualizado.ProductoID {
			productos[i] = productoActualizado
			return productos, nil
		}
	}
	return productos, errors.New("Producto no encontrado")
}

func RemoverProducto(productos []Producto, productoID int) ([]Producto, error) {
	for i, producto := range productos {
		if producto.ProductoID == productoID {
			return append(productos[:i], productos[i+1:]...), nil
		}
	}
	return productos, errors.New("Producto no encontrado")
}

func BuscarProductoNombre(productos []Producto, nombreProducto string) (Producto, error) {
	for _, producto := range productos {
		if producto.NombreProducto == nombreProducto {
			return producto, nil
		}
	}
	return Producto{}, errors.New("Producto no encontrado")
}

// Ejemplo de uso de concurrencia con goroutines y canales
func ProcesarProductosConcurrencia(productos []Producto) {
	ch := make(chan string, len(productos))

	for _, producto := range productos {
		go func(prod Producto) {
			ch <- fmt.Sprintf("Producto procesado: ID %d, Nombre %s", prod.ProductoID, prod.NombreProducto)
		}(producto)
	}

	for i := 0; i < len(productos); i++ {
		fmt.Println(<-ch)
	}
}
