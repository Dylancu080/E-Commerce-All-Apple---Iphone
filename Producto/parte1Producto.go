package producto

import (
	"errors"
	"fmt"
)

type Productos struct {
	productoID     int
	UserID         int
	nombreProducto string
	descripcion    string
	precio         float64
	categoria      string
	stok           int
}
type ListaProductos struct {
	lista []Productos
}

func agregarProducto(productos []Productos, nuevoProducto Productos) []Productos {
	return append(productos, nuevoProducto)
}

func (u *ListaProductos) actualizarProducto(nombreProducto, descripcion string, precio float64, stok int) {
	for i := range u.lista {
		if u.lista[i].nombreProducto == nombreProducto {
			u.lista[i].descripcion = descripcion
			u.lista[i].precio = precio
			u.lista[i].stok = stok
			return
		}
	}
	fmt.Println("Producto no encontrado:", nombreProducto)
}

func removerProducto(productos []Productos, productoID int) ([]Productos, error) {
	for i, producto := range productos {
		if producto.productoID == productoID {
			return append(productos[:i], productos[i+1:]...), nil
		}
	}
	return productos, errors.New("producto no encontrado")
}

func buscarProductoNombre(productos []Productos, nombreProducto string) (Productos, error) {
	for _, producto := range productos {
		if producto.nombreProducto == nombreProducto {
			return producto, nil
		}
	}
	return Productos{}, errors.New("El producto ingresado no se ha encontrado")
}
