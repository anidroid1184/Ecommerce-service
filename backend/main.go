package main

import (
    "ecommerce/backend/database"
    "ecommerce/backend/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Inicializar la base de datos
    database.InitDB()

    // Crear el servidor Gin
    server := gin.Default()

    // Definir las rutas
    routes.SetupRoutes(server)

    // Iniciar el servidor
    server.Run(":8080")
}
