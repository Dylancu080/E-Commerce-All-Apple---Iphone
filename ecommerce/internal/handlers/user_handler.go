package handlers

import (
	"ecommerce/internal/service"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	// Especificar la ruta absoluta de las plantillas HTML
	templates := template.Must(template.ParseGlob("C:/Programacion Orientada a Objetos/Autonomo 1a/ecommerce/web/templates/*.html"))

	err := templates.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		log.Println(err)
	}
}

func (h *UserHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	// Especificar la ruta absoluta de las plantillas HTML
	templates := template.Must(template.ParseGlob("C:/Programacion Orientada a Objetos/Autonomo 1a/ecommerce/web/templates/*.html"))

	err := templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Println(err)
	}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
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
	primerNombre := r.FormValue("primerNombre")
	segundoNombre := r.FormValue("segundoNombre")
	primerApellido := r.FormValue("primerApellido")
	segundoApellido := r.FormValue("segundoApellido")
	email := r.FormValue("email")
	password := r.FormValue("password")
	numero := r.FormValue("numero")
	genero := r.FormValue("genero")
	edadStr := r.FormValue("edad") // Obtener edad como string
	fechaNacimiento := r.FormValue("fechaNacimiento")
	esAdmin := r.FormValue("esAdmin") == "on"
	claveIngresada := r.FormValue("claveMaestra")

	// Convertir edad de string a int
	edad, err := strconv.Atoi(edadStr)
	if err != nil {
		http.Error(w, "Edad debe ser un número válido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para crear un nuevo usuario
	_, err = h.userService.CreateUser(primerNombre, segundoNombre, primerApellido, segundoApellido, email, password, numero, genero, edad, fechaNacimiento, esAdmin, claveIngresada)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirigir al login después del registro
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
