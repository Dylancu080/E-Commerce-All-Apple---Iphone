package handlers

import (
	"ecommerce/internal/service"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type PedidoHandler struct {
	pedidoService *service.PedidoService
}

func NewPedidoHandler(pedidoService *service.PedidoService) *PedidoHandler {
	return &PedidoHandler{pedidoService: pedidoService}
}

func (h *PedidoHandler) PedidoPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/pedido.html")
	if err != nil {
		http.Error(w, "Error al cargar la página de pedidos", http.StatusInternalServerError)
		log.Println("Error al cargar la página de pedidos:", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error al renderizar la página de pedidos", http.StatusInternalServerError)
		log.Println("Error al renderizar la página de pedidos:", err)
	}
}

func (h *PedidoHandler) AddPedido(w http.ResponseWriter, r *http.Request) {
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

	// Obtener datos del formulario
	usuarioIDStr := r.FormValue("usuarioID")

	// Convertir usuarioID de cadena a entero
	usuarioID, err := strconv.Atoi(usuarioIDStr)
	if err != nil {
		http.Error(w, "ID de usuario debe ser un número válido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para crear un nuevo pedido
	_, err = h.pedidoService.CreatePedido(usuarioID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirigir a la página de pedidos después de añadir un nuevo pedido
	http.Redirect(w, r, "/pedido", http.StatusSeeOther)
}
