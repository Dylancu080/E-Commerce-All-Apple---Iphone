//@autor: Dylan Curay
//@Fecha: 06 / 10 / 2024
//@Version: 1.0
//@Funcion: Separamos y definimos los tipos de entrega

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

	cities := []string{"Quito", "Provincia"} // 0 = 1 || 1 = 2

	e.deliveryCity = cities[cityOption-1]

	if e.deliveryCity == "Quito" {
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
		e.deliveryAddres = pongadirecciónProvincial()

	}

	fmt.Printf("Tipo de entrega: %s", e.tipoEntrega)

}
