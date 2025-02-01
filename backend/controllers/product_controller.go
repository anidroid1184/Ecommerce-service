package controllers

import (
    "net/http"
    "ecommerce-backend/models" // Importa el paquete models
    "github.com/gin-gonic/gin"
)

// Obtener todos los productos
func GetProducts(c *gin.Context) {
    var products []models.Product
    if err := models.DB.Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
        return
    }
    c.JSON(http.StatusOK, products)
}

// Crear un nuevo producto
func CreateProduct(c *gin.Context) {
    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := models.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }
    c.JSON(http.StatusCreated, input)
}