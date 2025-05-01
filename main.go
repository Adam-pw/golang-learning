package main

import (
    "github.com/gin-gonic/gin"
    "github.com/adam-pw/database"
    "github.com/adam-pw/routes"
)

func main() {
    r := gin.Default()
    database.Connect()      // connect to DB
    routes.RegisterRoutes(r) // register routes
    r.Run() // listen on :8080
}