package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// respuesta exitosa
func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"success": true,
		"data":    data,
	})
}

// respuesta de error
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"error":   message,
	})
}

// errores de validación
func ValidationErrorResponse(c *gin.Context, errors []string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"errors":  errors,
	})
}
