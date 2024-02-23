package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpportunityHandler(context *gin.Context) {
	request := CreateOpportunityRequest{}

	context.BindJSON(&request)

	logger.Infof("Request received: %+v", request)
}

func ShowOpportunityHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET opportunity",
	})
}

func UpdateOpportunityHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PUT opportunity",
	})
}

func DeleteOpportunityHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "DELETE opportunity",
	})
}

func ListOpportunitiesHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET opportunities",
	})
}
