package handlers

import (
	"ecommerce/internal/service"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ResenaHandler struct {
	resenaService *service.ResenaService
}

func NewResenaHandler(resenaService *service.ResenaService) *ResenaHandler {
	return &ResenaHandler{resenaService: resenaService}
}

func (h *ResenaHandler) ResenaPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/resena.html")
	if err != nil {
		http.Error(w, "Error al cargar la página de reseñas", http.StatusInternalServerError)
		log.Println("Error al cargar la página de reseñas:", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error al renderizar la página de reseñas", http.StatusInternalServerError)
		log.Println("Error al renderizar la página de reseñas:", err)
	}
}

func (h *ResenaHandler) AddResena(w http.ResponseWriter, r *http.Request) {
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
	productoIDStr := r.FormValue("productoID")
	calificacionStr := r.FormValue("calificacion")
	comentario := r.FormValue("comentario")

	// Convertir productoID y usuarioID de cadena a entero
	productoID, err := strconv.Atoi(productoIDStr)
	if err != nil {
		http.Error(w, "ID de producto debe ser un número válido", http.StatusBadRequest)
		return
	}

	// Convertir calificacion de string a int
	calificacion, err := strconv.Atoi(calificacionStr)
	if err != nil {
		http.Error(w, "Calificación debe ser un número válido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para crear una nueva reseña
	_, err = h.resenaService.CreateResena(productoID, calificacion, comentario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirigir a la página de reseñas después de añadir una nueva reseña
	http.Redirect(w, r, "/resena", http.StatusSeeOther)
}
