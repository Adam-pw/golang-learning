package main

import (
    "log"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "github.com/adam-pw/database"
    "github.com/adam-pw/routes"
)

func main() {
    err := godotenv.Load() // load .env file
    if err != nil {
        log.Fatal("Error loading .env file") // log error if .env file is not found
    }
    r := gin.Default()
    database.Connect()      // connect to DB
    routes.RegisterRoutes(r) // register routes
    r.Run() // listen on :8080
}