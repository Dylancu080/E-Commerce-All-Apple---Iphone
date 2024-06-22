package entregapedido

import "fmt"

type Entregas struct {
	pedidoID        int
	UserID          int
	precioTotal     int
	status          string
	tipoEntrega     string
	deliveryCity    string
	deliveryAddres  string
	meetingLocation string
}

func EscogerTipoEntrega(e *Entregas, cityOption int) {
	cities := []string{"Quito", "Provincia"}
	e.deliveryCity = cities[cityOption-1]