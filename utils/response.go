package utils

import (
	"github.com/gin-gonic/gin"
)

// Response is the standard API response structure
type Response struct {
	IsSuccess  bool        `json:"is_success"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Meta       interface{} `json:"meta"`
}

// SuccessResponse returns a standardized success response
func SuccessResponse(c *gin.Context, statusCode int, data interface{}, meta interface{}) {
	response := Response{
		IsSuccess:  true,
		StatusCode: statusCode,
		Data:       data,
		Meta:       meta,
	}
	
	c.JSON(statusCode, response)
}

// ErrorResponse returns a standardized error response
func ErrorResponse(c *gin.Context, statusCode int, err string) {
	response := Response{
		IsSuccess:  false,
		StatusCode: statusCode,
		Data:       gin.H{"error": err},
		Meta:       nil,
	}
	
	c.JSON(statusCode, response)
} 