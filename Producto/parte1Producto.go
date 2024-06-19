package producto

import "errors"

type Productos struct {
	UserID         int
	nombreProducto string
	descripcion    string
	precio         float64
	categoria      string
	stok           int
}

func agregarProducto(productos []Productos, nuevoProducto Productos) []Productos {
	return append(productos, nuevoProducto)
}

func (u *Productos) actualizarProducto(nombreProducto, descripcion string, precio float64, stok int) {
	u.nombreProducto = nombreProducto
	u.descripcion = descripcion
	u.precio = precio
	u.stok = stok
}

func removerProducto(productos []Productos, id int) ([]Productos, error) {
	for i, producto := range productos {
		if producto.ID == id {
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
