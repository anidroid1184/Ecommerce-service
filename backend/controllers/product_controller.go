package controllers

import (
    "net/http"
    "ecommerce/backend/models"
    "github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
    var products []models.Product
    models.DB.Find(&products)
    c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Create(&input)
    c.JSON(http.StatusCreated, input)
}
