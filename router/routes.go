package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rainanDeveloper/gopportunities/docs"
	"github.com/rainanDeveloper/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api"
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Title = "GOpportunity"
	api := router.Group(basePath)
	{
		api.GET("/opportunity/:id", handler.ShowOpportunityHandler)
		api.POST("/opportunity", handler.CreateOpportunityHandler)
		api.PUT("/opportunity/:id", handler.UpdateOpportunityHandler)
		api.DELETE("/opportunity/:id", handler.DeleteOpportunityHandler)
		api.GET("/opportunities", handler.ListOpportunitiesHandler)
	}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
