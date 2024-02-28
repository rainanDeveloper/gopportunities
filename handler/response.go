package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rainanDeveloper/gopportunities/schemas"
)

func sendResponseError(context *gin.Context, code int, message string) {
	context.Header("Content-Type", "application/json")
	context.JSON(code, gin.H{
		"message":     message,
		"status_code": code,
	})
}

func sendResponseSuccess(context *gin.Context, data interface{}) {
	context.Header("Content-Type", gin.MIMEJSON)
	context.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func sendResponseNoContent(context *gin.Context) {
	context.Data(http.StatusNoContent, gin.MIMEHTML, nil)
}

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode string `json:"status_code"`
}

type CreateOpportunityResponse struct {
	Data schemas.OpportunityResponseBody `json:"data"`
}
