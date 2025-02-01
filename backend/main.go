package main

import (
	"ecommerce-backend/database"
	"ecommerce-backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la base de datos
	database.InitDB()

	// Crear el servidor Gin
	r := gin.Default()

	// habilitar CORS
	r.Use(cors.Default())

	// Definir las rutas
	routes.SetupRoutes(r)

	// Iniciar el servidor
	r.Run(":8080")

}
