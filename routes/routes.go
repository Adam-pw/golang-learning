package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/adam-pw/controllers"
)

func RegisterRoutes(r *gin.Engine) {
    // User Routes
    r.POST("/users", controllers.CreateUser)
    r.GET("/users/:id", controllers.GetUser)
    r.GET("/users", controllers.GetAllUsers)
    r.PUT("/users/:id", controllers.UpdateUser)
    r.DELETE("/users/:id", controllers.DeleteUser)
}
