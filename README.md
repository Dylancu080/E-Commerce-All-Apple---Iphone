# E-Commerce All Apple - iPhone

**Fecha**: 29 de junio de 2024

**Integrantes**:
- Dylan Curay

## Descripción
Este proyecto de e-commerce está desarrollado en Go y se enfoca en la venta de productos Apple, específicamente iPhones. 
El sistema permite a los usuarios registrarse, iniciar sesión, añadir productos al carrito, realizar pedidos, y gestionar entregas y detalles de pedidos.

## Funcionalidades Principales
- Registro e inicio de sesión de usuarios.
- Visualización y compra de iPhones.
- Creación y gestión de pedidos.
- Gestión de entregas y detalles de pedidos.

## Objetivo del Programa
El objetivo de este programa es implementar un sistema de e-commerce funcional utilizando Go, 
demostrando el uso de bases de datos y servicios web para gestionar un catálogo de productos y pedidos de manera eficiente.


## Principales Funcionalidades del Codigo

main.go

----------------------------------------------------------------------------------------------------------------------------------------------

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
                                                                                                        //main, como su nombre lo indica, es la parte primordial para poder
                                                                                                        //importar todas las clases necesarias e imprimir nuestro codigo
                                                                                                      
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


-------------------------------------------------------------------------------------------------------------------------------------------------------------

Manejo de Repositorios

package repository

import (
	"database/sql"
	"ecommerce/internal/models"
)
                                                                                                      //Es muy importante esta fase ya que implementamos 
                                                                                                      //SQL en nuestro codigo y es donde se va guardar todos 
                                                                                                      //nuestros datos
type CategoriaRepository struct {
	db *sql.DB
}

func NewCategoriaRepository(db *sql.DB) *CategoriaRepository {
	return &CategoriaRepository{db: db}
}

func (repo *CategoriaRepository) CreateCategoria(categoria *models.Categoria) error {
	query := "INSERT INTO categorias (nombre, descripcion) VALUES (?, ?)"
	_, err := repo.db.Exec(query, categoria.NombreCategoria, categoria.Descripcion)
	return err
}

------------------------------------------------------------------------------------------------------------------------------------------------------------

categoria.html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">                        
    <title>Agregar Categoría</title>
</head>                                                                                              //Usamos HTML para poder implementar nuestro codigo en el front END
                                                                                                     //en el back END y todo, aqui tambien usamos CSS para diseño amigable
                                                                                                     //con el usuario
<body>
    <h1>Agregar Categoría</h1>
    <form action="/categoria" method="post">
        <label for="nombre">Nombre:</label>
        <input type="text" id="nombre" name="nombre"><br>

        <label for="descripcion">Descripción:</label>
        <input type="text" id="descripcion" name="descripcion"><br>

        <input type="submit" value="Agregar Categoría">
    </form>
</body>
</html>
----------------------------------------------------------------------------------------------------------------------------------------------------------


//Si deseas ejecutar el programa, usa "go run main.go"

Gracias
