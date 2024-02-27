package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rainanDeveloper/gopportunities/schemas"
)

func CreateOpportunityHandler(context *gin.Context) {
	request := CreateOpportunityRequest{}

	context.BindJSON(&request)

	oppotunity := schemas.Opportunity{
		ID:        uuid.New(),
		Role:      request.Role,
		Company:   request.Company,
		Location:  request.Location,
		Remote:    *request.Remote,
		Link:      request.Link,
		Salary:    request.Salary,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&oppotunity).Error; err != nil {
		logger.Errorf("Error creating opportunity: %v", err.Error())
		sendResponseError(context, http.StatusInternalServerError, "Error during opportunity creation on database")
		return
	}

	sendResponseSuccess(context, oppotunity)
	return
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
