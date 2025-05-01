package database

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/adam-pw/models"
)

var DB *gorm.DB

func Connect() {
    dsn := "host=localhost user=adampithenwala password=12345678 dbname=go_gin_db port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    fmt.Println("Database connection established")

    // Auto-migrate your models here
    DB.AutoMigrate(&models.User{})
}
