package main

import (
	"database/sql"
	"ecommerce/internal/handlers"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb" // Importación del controlador SQL Server
	"github.com/gorilla/mux"
)

func main() {
	// Crear conexión a la base de datos
	connString := "sqlserver://localhost?database=DYLAN&trusted_connection=yes"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Repositorios
	userRepo := repository.NewUserRepository(db)
	categoriaRepo := repository.NewCategoriaRepository(db)
	detallesPedidoRepo := repository.NewDetallesPedidoRepository(db)
	entregaPedidoRepo := repository.NewEntregaPedidoRepository(db)
	pedidoRepo := repository.NewPedidoRepository(db)
	productoRepo := repository.NewProductoRepository(db)
	resenaRepo := repository.NewResenaRepository(db)

	// Servicios
	userService := service.NewUserService(userRepo)
	categoriaService := service.NewCategoriaService(categoriaRepo)
	detallesPedidoService := service.NewDetallesPedidoService(detallesPedidoRepo)
	entregaPedidoService := service.NewEntregaPedidoService(entregaPedidoRepo)
	pedidoService := service.NewPedidoService(pedidoRepo)
	productoService := service.NewProductoService(productoRepo)
	resenaService := service.NewResenaService(resenaRepo)

	// Manejadores
	userHandler := handlers.NewUserHandler(userService)
	categoriaHandler := handlers.NewCategoriaHandler(categoriaService)
	detallesPedidoHandler := handlers.NewDetallesPedidoHandler(detallesPedidoService)
	entregaPedidoHandler := handlers.NewEntregaPedidoHandler(entregaPedidoService)
	pedidoHandler := handlers.NewPedidoHandler(pedidoService)
	productoHandler := handlers.NewProductoHandler(productoService)
	resenaHandler := handlers.NewResenaHandler(resenaService)

	// Rutas
	r := mux.NewRouter()

	// Rutas de usuarios
	r.HandleFunc("/register", userHandler.RegisterPage).Methods("GET")
	r.HandleFunc("/login", userHandler.LoginPage).Methods("GET")
	r.HandleFunc("/registerUser", userHandler.RegisterUser).Methods("POST")

	// Rutas de categorías
	r.HandleFunc("/categoria", categoriaHandler.CategoriaPage).Methods("GET")
	r.HandleFunc("/addCategoria", categoriaHandler.AddCategoria).Methods("POST")

	// Rutas de productos
	r.HandleFunc("/producto", productoHandler.ProductoPage).Methods("GET")
	r.HandleFunc("/addProducto", productoHandler.AddProducto).Methods("POST")

	// Rutas de pedidos
	r.HandleFunc("/pedido", pedidoHandler.PedidoPage).Methods("GET")
	r.HandleFunc("/addPedido", pedidoHandler.AddPedido).Methods("POST")

	// Rutas de detalles de pedidos
	r.HandleFunc("/detallesPedido", detallesPedidoHandler.DetallesPedidoPage).Methods("GET")
	r.HandleFunc("/addDetallesPedido", detallesPedidoHandler.AddDetallesPedido).Methods("POST")

	// Rutas de entregas de pedidos
	r.HandleFunc("/entregaPedido", entregaPedidoHandler.EntregaPedidoPage).Methods("GET")
	r.HandleFunc("/addEntregaPedido", entregaPedidoHandler.AddEntregaPedido).Methods("POST")

	// Rutas de reseñas
	r.HandleFunc("/resena", resenaHandler.ResenaPage).Methods("GET")
	r.HandleFunc("/addResena", resenaHandler.AddResena).Methods("POST")

	// Servir archivos estáticos
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	// Iniciar servidor
	http.ListenAndServe(":8080", r)
}
