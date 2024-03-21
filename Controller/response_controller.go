package Controller

import (
	"encoding/json"
	"net/http"

	m "FrameworkAPI/Model"

	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ErrorResponse
	response.Status = code
	response.Message = message

	json.NewEncoder(w).Encode(response)
}

func SendErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ErrorResponse
	response.Status = code
	response.Message = message

	json.NewEncoder(w).Encode(response)
}

func SendSuccessResponseGIN(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

//	func SendErrorResponseGIN(c *gin.Context, code int, message string) {
//		c.JSON(code, gin.H{
//			"code":  code,
//			"error": message,
//			// "data":  product,
//		})
//	}
func SendErrorResponseGIN(c *gin.Context, code int, message string) {
	errResponse := m.ErrorResponse{
		Status:  code,
		Message: message,
	}
	c.JSON(code, errResponse)
}

func SendDataResponseGIN(c *gin.Context, code int, message string, products []m.Products) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    products,
	})
}
