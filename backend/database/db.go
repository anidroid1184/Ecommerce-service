package database

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrar el esquema
    DB.AutoMigrate(&Product{})
    fmt.Println("Database connected and migrated!")
}
