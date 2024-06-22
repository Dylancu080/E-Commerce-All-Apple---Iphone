package categoria

type Categoria struct {
	ID     int
	Nombre string
}

func agregar(c []Categoria, n Categoria) []Categoria {
	return append(c, n)
}

func remover(c []Categoria, id int) []Categoria {
	for i, cat := range c {
		if cat.ID == id {
			return append(c[:i], c[i+1:]...)
		}
	}
	return c
}
