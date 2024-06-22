package producto

type Productos struct {
	ID          int
	UserID      int
	Nombre      string
	Descripcion string
	Precio      float64
	Categoria   string
	Stock       int
}

func agregar(productos []Productos, nuevoProducto Productos) []Productos {
	return append(productos, nuevoProducto)
}

func actualizar(productos []Productos, nombre, descripcion string, precio float64, stock int) []Productos {
	for i := range productos {
		if productos[i].Nombre == nombre {
			productos[i].Descripcion = descripcion
			productos[i].Precio = precio
			productos[i].Stock = stock
			break
		}
	}
	return productos
}
