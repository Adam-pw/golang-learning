package controllers

import (
    "net/http"
    "github.com/adam-pw/database"
    "github.com/adam-pw/models"
    "github.com/adam-pw/utils"
    "github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
    var user models.User

    // Bind JSON to user struct
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    // Validate required fields
    if user.Email == "" {
        utils.ErrorResponse(c, http.StatusBadRequest, "email is required")
        return
    }

    // Insert into database
    result := database.DB.Create(&user)
    if result.Error != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, result.Error.Error())
        return
    }

    // Respond with created user
    utils.SuccessResponse(c, http.StatusOK, user, nil)
}

func GetUser(c *gin.Context) {
    // Get id from url parameter
    id := c.Param("id")
    
    var user models.User
    
    // Find the user in database
    result := database.DB.First(&user, id)
    if result.Error != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "User not found")
        return
    }
    
    // Return the user
    utils.SuccessResponse(c, http.StatusOK, user, nil)
}

func GetAllUsers(c *gin.Context) {
    var users []models.User
    
    // Get all users from database
    result := database.DB.Find(&users)
    if result.Error != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, result.Error.Error())
        return
    }
    
    // Return the users
    utils.SuccessResponse(c, http.StatusOK, users, nil)
}

func UpdateUser(c *gin.Context) {
    // Get id from url parameter
    id := c.Param("id")
    
    // Check if user exists
    var existingUser models.User
    if result := database.DB.First(&existingUser, id); result.Error != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "User not found")
        return
    }
    
    // Bind JSON to input struct
    var input struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    
    // Update user fields if provided
    updateData := map[string]interface{}{}
    
    if input.Name != "" {
        updateData["name"] = input.Name
    }
    
    if input.Email != "" {
        updateData["email"] = input.Email
    }
    
    // Update the user in database
    if result := database.DB.Model(&existingUser).Updates(updateData); result.Error != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, result.Error.Error())
        return
    }
    
    // Get updated user
    database.DB.First(&existingUser, id)
    
    // Return the updated user
    utils.SuccessResponse(c, http.StatusOK, existingUser, nil)
}

func DeleteUser(c *gin.Context) {
    // Get id from url parameter
    id := c.Param("id")
    
    // Check if user exists
    var user models.User
    if result := database.DB.First(&user, id); result.Error != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "User not found")
        return
    }
    
    // Delete user from database
    if result := database.DB.Delete(&user); result.Error != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, result.Error.Error())
        return
    }
    
    // Return success response
    utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "User deleted successfully"}, nil)
}
