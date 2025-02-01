package database

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "ecommerce-backend/models" // Importa el paquete models
)

// Función para inicializar la base de datos
func InitDB() {
    var err error
    models.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrar el esquema
    models.DB.AutoMigrate(&models.Product{})
    fmt.Println("Database connected and migrated!")
}