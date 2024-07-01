package handlers

import (
	"ecommerce/internal/service"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type DetallesPedidoHandler struct {
	detallesPedidoService *service.DetallesPedidoService
}

func NewDetallesPedidoHandler(detallesPedidoService *service.DetallesPedidoService) *DetallesPedidoHandler {
	return &DetallesPedidoHandler{detallesPedidoService: detallesPedidoService}
}

func (h *DetallesPedidoHandler) DetallesPedidoPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/detalles_pedido.html")
	if err != nil {
		http.Error(w, "Error al cargar la página de detalles de pedido", http.StatusInternalServerError)
		log.Println("Error al cargar la página de detalles de pedido:", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error al renderizar la página de detalles de pedido", http.StatusInternalServerError)
		log.Println("Error al renderizar la página de detalles de pedido:", err)
	}
}

func (h *DetallesPedidoHandler) AddDetallesPedido(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Parsear el formulario
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al parsear el formulario", http.StatusBadRequest)
		return
	}

	// Obtener datos del formulario y convertir a int
	pedidoID, err := strconv.Atoi(r.FormValue("pedidoID"))
	if err != nil {
		http.Error(w, "ID de pedido inválido", http.StatusBadRequest)
		return
	}

	productoID, err := strconv.Atoi(r.FormValue("productoID"))
	if err != nil {
		http.Error(w, "ID de producto inválido", http.StatusBadRequest)
		return
	}

	cantidad, err := strconv.Atoi(r.FormValue("cantidad"))
	if err != nil {
		http.Error(w, "Cantidad inválida", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para crear un nuevo detalle de pedido
	_, err = h.detallesPedidoService.CreateDetallesPedido(pedidoID, productoID, cantidad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirigir a la página de detalles de pedidos después de añadir un nuevo detalle de pedido
	http.Redirect(w, r, "/detallesPedido", http.StatusSeeOther)
}
