//@autor: Dylan Curay
//@Fecha: 06 / 10 / 2024
//@Version: 1.0
//@Funcion: Separamos y definimos los tipos de entrega

package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entrega struct {
	PedidoID        int
	UserID          int
	PrecioTotal     float64
	Status          string
	TipoEntrega     string
	DeliveryCity    string
	DeliveryAddress string
	MeetingLocation string
}

type EntregaInterface interface {
	EscogerTipoEntrega(cityOption int)
	ObtenerUbicacionEncuentro() string
	ObtenerDireccionEntrega() string
	ConsultarEstado() string
	ActualizarEstado(newStatus string)
}

func (e *Entrega) EscogerTipoEntrega(cityOption int) {
	cities := []string{"Quito", "Provincia"}

	e.DeliveryCity = cities[cityOption-1]

	if e.DeliveryCity == "Quito" {
		e.TipoEntrega = "Presencial"
		meetingLocations := []string{"El condado Shopping", "Quicentro Shopping", "CCI"}
		escogerLocation := seleccionarMeetingLocation(meetingLocations)
		if escogerLocation != "" {
			e.MeetingLocation = escogerLocation
		} else {
			e.MeetingLocation = ingresarUbicacionPersonalizada()
		}
	} else {
		e.TipoEntrega = "Envío"
		e.DeliveryAddress = ingresarDireccionProvincial()
	}

	fmt.Printf("Tipo de entrega: %s\n", e.TipoEntrega)
}

func seleccionarMeetingLocation(locations []string) string {
	fmt.Println("Selecciona la ubicación de encuentro en Quito:")
	for i, loc := range locations {
		fmt.Printf("%d: %s\n", i+1, loc)
	}
	fmt.Println("0: Ingresar otra ubicación")

	var option int
	fmt.Scanln(&option)

	if option == 0 {
		return ""
	}

	if option < 1 || option > len(locations) {
		fmt.Println("Opción inválida")
		return seleccionarMeetingLocation(locations)
	}

	return locations[option-1]
}

func ingresarUbicacionPersonalizada() string {
	fmt.Println("Ingresa la ubicación de encuentro personalizada en Quito:")
	reader := bufio.NewReader(os.Stdin)
	customLocation, _ := reader.ReadString('\n')
	return strings.TrimSpace(customLocation)
}

func ingresarDireccionProvincial() string {
	fmt.Println("Ingresa la dirección de envío:")
	reader := bufio.NewReader(os.Stdin)
	deliveryAddress, _ := reader.ReadString('\n')
	return strings.TrimSpace(deliveryAddress)
}

func (e *Entrega) ObtenerUbicacionEncuentro() string {
	if e.TipoEntrega == "Presencial" {
		return e.MeetingLocation
	}
	return ""
}

func (e *Entrega) ObtenerDireccionEntrega() string {
	if e.TipoEntrega == "Envío" {
		return e.DeliveryAddress
	}
	return ""
}

func (e *Entrega) ConsultarEstado() string {
	return e.Status
}

func (e *Entrega) ActualizarEstado(newStatus string) {
	e.Status = newStatus
	fmt.Printf("Estado del pedido actualizado a: %s\n", e.Status)
}
