package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/opportunity", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET opportunity",
			})
		})
		api.POST("/opportunity", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "POST opportunity",
			})
		})
		api.PUT("/opportunity", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "PUT opportunity",
			})
		})
		api.DELETE("/opportunity", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "DELETE opportunity",
			})
		})
		api.GET("/opportunities", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET opportunities",
			})
		})
	}
}
