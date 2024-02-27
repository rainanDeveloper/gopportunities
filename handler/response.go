package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendResponseError(context *gin.Context, code int, message string) {
	context.Header("Content-Type", "application/json")
	context.JSON(code, gin.H{
		"message":    message,
		"statusCode": code,
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
