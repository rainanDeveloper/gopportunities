package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rainanDeveloper/gopportunities/schemas"
)

func CreateOpportunityHandler(context *gin.Context) {
	request := CreateOpportunityRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendResponseError(context, http.StatusBadRequest, err.Error())
		return
	}

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
		logger.Errorf("error creating opportunity: %v", err.Error())
		sendResponseError(context, http.StatusInternalServerError, "error during opportunity creation on database")
		return
	}

	sendResponseSuccess(context, oppotunity)
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
	id := context.Param("id")

	opportunity := schemas.Opportunity{}

	if err := db.First(&opportunity, "id = ?", id).Error; err != nil {
		sendResponseError(context, http.StatusNotFound, fmt.Sprintf("opportunity with id %s not found", id))
		return
	}

	if err := db.Delete(&opportunity).Error; err != nil {
		logger.Errorf("error deleting opportunity with id %s", id)
		sendResponseError(context, http.StatusInternalServerError, fmt.Sprintf("error deleting opportunity with id %s", id))
		return
	}

	sendResponseNoContent(context)
}

func ListOpportunitiesHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET opportunities",
	})
}
