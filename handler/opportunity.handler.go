package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rainanDeveloper/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Create Opportunity
// @Description Create a new job opportunity
// @Tags Opportunities
// @Accept json
// @Produce json
// @Param request body CreateOpportunityRequest true "Request body"
// @Success 200 {object} CreateOpportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunity [post]
func CreateOpportunityHandler(context *gin.Context) {
	request := CreateOpportunityRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendResponseError(context, http.StatusBadRequest, err.Error())
		return
	}

	opportunity := schemas.Opportunity{
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

	if err := db.Create(&opportunity).Error; err != nil {
		logger.Errorf("error creating opportunity: %v", err.Error())
		sendResponseError(context, http.StatusInternalServerError, "error during opportunity creation on database")
		return
	}

	responseData := schemas.OpportunityResponseBody(opportunity)

	sendResponseSuccess(context, responseData)
}

// @Summary Find a Opportunity
// @Description Find a specific job opportunity by id
// @Tags Opportunities
// @Accept json
// @Produce json
// @Param id path string true "Opportunity Id"
// @Success 200 {object} ShowOpportunityResponse
// @Failure 400 {object} ErrorResponse
// @Router /opportunity/{id} [get]
func ShowOpportunityHandler(context *gin.Context) {
	id := context.Param("id")

	opportunity := schemas.Opportunity{}

	if err := db.First(&opportunity, "id = ?", id).Error; err != nil {
		sendResponseError(context, http.StatusNotFound, fmt.Sprintf("opportunity with id %s not found", id))
		return
	}

	responseData := schemas.OpportunityResponseBody(opportunity)

	sendResponseSuccess(context, responseData)
}

// @Summary Update a Opportunity
// @Description Update a specific job opportunity by id
// @Tags Opportunities
// @Accept json
// @Produce json
// @Param id path string true "Opportunity Id"
// @Param request body UpdateOpportunityRequest true "Request body"
// @Success 200 {object} UpdateOpportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunity/{id} [put]
func UpdateOpportunityHandler(context *gin.Context) {
	request := UpdateOpportunityRequest{}

	context.BindJSON(&request)
	id := context.Param("id")

	if err := request.Validate(); err != nil {
		sendResponseError(context, http.StatusBadRequest, err.Error())
		return
	}

	opportunity := schemas.Opportunity{}

	if err := db.First(&opportunity, "id = ?", id).Error; err != nil {
		sendResponseError(context, http.StatusNotFound, fmt.Sprintf("opportunity with id %s not found", id))
		return
	}

	if request.Role != "" {
		opportunity.Role = request.Role
	}

	if request.Company != "" {
		opportunity.Company = request.Company
	}

	if request.Link != "" {
		opportunity.Link = request.Link
	}

	if request.Location != "" {
		opportunity.Location = request.Location
	}

	if request.Remote != nil {
		opportunity.Remote = *request.Remote
	}

	if request.Salary > 0 {
		opportunity.Salary = request.Salary
	}

	if err := db.Save(&opportunity).Error; err != nil {
		sendResponseError(context, http.StatusInternalServerError, "error while updating the opportunity on database")
		return
	}

	responseData := schemas.OpportunityResponseBody(opportunity)

	sendResponseSuccess(context, responseData)
}

// @Summary Delete a Opportunity
// @Description Delete a specific job opportunity by id
// @Tags Opportunities
// @Accept json
// @Produce json
// @Param id path string true "Opportunity Id"
// @Success 204 {object} nil
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunity/{id} [delete]
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

// @Summary List Opportunities
// @Description List all job opportunities available
// @Tags Opportunities
// @Accept json
// @Produce json
// @Success 200 {object} ListOpportunitiesResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunities [get]
func ListOpportunitiesHandler(context *gin.Context) {
	opportunities := []schemas.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		logger.Errorf("error listing opportunities")
		sendResponseError(context, http.StatusInternalServerError, "error listing opportunities")
		return
	}

	responseData := []schemas.OpportunityResponseBody{}

	for i := 0; i < len(opportunities); i++ {
		responseData = append(responseData, schemas.OpportunityResponseBody(opportunities[i]))
	}

	sendResponseSuccess(context, responseData)
}
