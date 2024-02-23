package router

import (
	"github.com/gin-gonic/gin"
	opportunity_handler "github.com/rainanDeveloper/gopportunities/handler/opportunity"
)

func initializeRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/opportunity_handler", opportunity_handler.ShowOpportunityHandler)
		api.POST("/opportunity_handler", opportunity_handler.CreateOpportunityHandler)
		api.PUT("/opportunity_handler", opportunity_handler.UpdateOpportunityHandler)
		api.DELETE("/opportunity_handler", opportunity_handler.DeleteOpportunityHandler)
		api.GET("/opportunities", opportunity_handler.ListOpportunitiesHandler)
	}
}
