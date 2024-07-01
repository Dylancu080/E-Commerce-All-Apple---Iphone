package handlers

import (
	"ecommerce/internal/service"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type EntregaPedidoHandler struct {
	entregaPedidoService *service.EntregaPedidoService
}

func NewEntregaPedidoHandler(entregaPedidoService *service.EntregaPedidoService) *EntregaPedidoHandler {
	return &EntregaPedidoHandler{entregaPedidoService: entregaPedidoService}
}

func (h *EntregaPedidoHandler) EntregaPedidoPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/entrega_pedido.html")
	if err != nil {
		http.Error(w, "Error al cargar la página de entrega de pedido", http.StatusInternalServerError)
		log.Println("Error al cargar la página de entrega de pedido:", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error al renderizar la página de entrega de pedido", http.StatusInternalServerError)
		log.Println("Error al renderizar la página de entrega de pedido:", err)
	}
}

func (h *EntregaPedidoHandler) AddEntregaPedido(w http.ResponseWriter, r *http.Request) {
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

	// Obtener datos del formulario y convertir pedidoID a int
	pedidoID, err := strconv.Atoi(r.FormValue("pedidoID"))
	if err != nil {
		http.Error(w, "ID de pedido inválido", http.StatusBadRequest)
		return
	}

	direccionEntrega := r.FormValue("direccionEntrega")

	// Llamar al servicio para crear una nueva entrega de pedido
	_, err = h.entregaPedidoService.CreateEntregaPedido(pedidoID, direccionEntrega)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirigir a la página de entrega de pedidos después de añadir una nueva entrega de pedido
	http.Redirect(w, r, "/entregaPedido", http.StatusSeeOther)
}
