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

func (e *Entregas) EscogerTipoEntrega(cityOption int) {

	cities := []string{"Quito", "Provincia"}

	e.deliveryCity = cities[cityOption-1]

	if BadExpr {
		e.tipoEntrega = "Presencial"

		meetingLocation := []string{"El condado Shopping", "Quicentro Shopping", "CCI"}
		escogerLocation := seleccionarMeetingLocation(meetingLocation)

		if escogerLocation != "" {
			e.meetingLocation = escogerLocation
		} else {
			e.meetingLocation = seleccionarMeetingLocation()
		}
	} else {
		e.tipoEntrega = "Envio"
		e.deliveryAddres = pongadirecciónProvincia()

	}

	fmt.Printf("Tipo de entrega: %s", e.tipoEntrega)

}
