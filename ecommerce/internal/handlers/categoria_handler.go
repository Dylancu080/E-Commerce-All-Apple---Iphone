package handlers

import (
	"ecommerce/internal/service"
	"html/template"
	"log"
	"net/http"
)

type CategoriaHandler struct {
	categoriaService *service.CategoriaService
}

func NewCategoriaHandler(categoriaService *service.CategoriaService) *CategoriaHandler {
	return &CategoriaHandler{categoriaService: categoriaService}
}

func (h *CategoriaHandler) CategoriaPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/categoria.html")
	if err != nil {
		http.Error(w, "Error al cargar la página de categoría", http.StatusInternalServerError)
		log.Println("Error al cargar la página de categoría:", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error al renderizar la página de categoría", http.StatusInternalServerError)
		log.Println("Error al renderizar la página de categoría:", err)
	}
}

func (h *CategoriaHandler) AddCategoria(w http.ResponseWriter, r *http.Request) {
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
	nombre := r.FormValue("nombre")
	descripcion := r.FormValue("descripcion")

	// Llamar al servicio para crear una nueva categoría
	_, err = h.categoriaService.CreateCategoria(nombre, descripcion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirigir a la página de categorías después de añadir una nueva categoría
	http.Redirect(w, r, "/categoria", http.StatusSeeOther)
}
