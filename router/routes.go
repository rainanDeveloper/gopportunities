package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rainanDeveloper/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	api := router.Group("/api")
	{
		api.GET("/opportunity/:id", handler.ShowOpportunityHandler)
		api.POST("/opportunity", handler.CreateOpportunityHandler)
		api.PUT("/opportunity/:id", handler.UpdateOpportunityHandler)
		api.DELETE("/opportunity/:id", handler.DeleteOpportunityHandler)
		api.GET("/opportunities", handler.ListOpportunitiesHandler)
	}
}
