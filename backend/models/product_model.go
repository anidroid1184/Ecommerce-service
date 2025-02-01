package models

import "gorm.io/gorm"

// Definición del modelo de Producto
type Product struct {
    gorm.Model
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

// Variable global para la conexión a la base de datos
var DB *gorm.DB // Asegúrate de que esta línea esté presente