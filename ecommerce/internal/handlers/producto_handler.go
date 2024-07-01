package handlers

import (
	"ecommerce/internal/service"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ProductoHandler struct {
	productoService *service.ProductoService
}

func NewProductoHandler(productoService *service.ProductoService) *ProductoHandler {
	return &ProductoHandler{productoService: productoService}
}

func (h *ProductoHandler) ProductoPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/producto.html")
	if err != nil {
		http.Error(w, "Error al cargar la página de productos", http.StatusInternalServerError)
		log.Println("Error al cargar la página de productos:", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error al renderizar la página de productos", http.StatusInternalServerError)
		log.Println("Error al renderizar la página de productos:", err)
	}
}

func (h *ProductoHandler) AddProducto(w http.ResponseWriter, r *http.Request) {
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
	precioStr := r.FormValue("precio")
	stockStr := r.FormValue("stock")

	// Convertir precio de string a float64
	precio, err := strconv.ParseFloat(precioStr, 64)
	if err != nil {
		http.Error(w, "Precio debe ser un número válido", http.StatusBadRequest)
		return
	}

	// Convertir stock de string a int
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		http.Error(w, "Stock debe ser un número válido", http.StatusBadRequest)
		return
	}

	// Supongamos que el ID del producto se genera automáticamente o se obtiene de otro lugar
	// En este caso, lo dejaremos como un valor fijo para fines de demostración
	productID := 1

	// Llamar al servicio para crear un nuevo producto
	_, err = h.productoService.CreateProducto(productID, nombre, descripcion, precio, stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirigir a la página de productos después de añadir un nuevo producto
	http.Redirect(w, r, "/producto", http.StatusSeeOther)
}
