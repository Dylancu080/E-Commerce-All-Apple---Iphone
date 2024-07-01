package models

import (
	"errors"
	"fmt"
	"sync"
)

type Detalles struct {
	DetalleID  int
	PedidoID   int
	ProductoID int
	Cantidad   int
	Precio     float64
}

// Interfaz para la estructura Detalles (opcional según necesidad)
type DetallesInterface interface {
	PagarOrden()
}

// Método para pagar una orden
func (d *Detalles) PagarOrden() {
	total := d.Precio * float64(d.Cantidad)
	fmt.Printf("Orden pagada. Detalle ID: %d, Total: $%.2f\n", d.DetalleID, total)
}

// Función para agregar un detalle a la lista de detalles
func AgregarDetalle(detalles []Detalles, nuevoDetalle Detalles) []Detalles {
	return append(detalles, nuevoDetalle)
}

// Función para actualizar un detalle en la lista de detalles
func ActualizarDetalle(detalles []Detalles, detalleActualizado Detalles) ([]Detalles, error) {
	for i, detalle := range detalles {
		if detalle.DetalleID == detalleActualizado.DetalleID {
			detalles[i] = detalleActualizado
			return detalles, nil
		}
	}
	return detalles, errors.New("Detalle no encontrado")
}

// Función para remover un detalle de la lista de detalles por su ID
func RemoverDetalle(detalles []Detalles, detalleID int) ([]Detalles, error) {
	for i, detalle := range detalles {
		if detalle.DetalleID == detalleID {
			return append(detalles[:i], detalles[i+1:]...), nil
		}
	}
	return detalles, errors.New("Detalle no encontrado")
}

// Ejemplo de procesamiento concurrente de detalles
func ProcesarDetallesConcurrencia(detalles []Detalles) {
	var wg sync.WaitGroup
	ch := make(chan string, len(detalles))

	for _, detalle := range detalles {
		wg.Add(1)
		go func(det Detalles) {
			defer wg.Done()
			det.PagarOrden()
			ch <- fmt.Sprintf("Detalle ID %d procesado", det.DetalleID)
		}(detalle)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
