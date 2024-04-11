package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gojobs/schemas"
)

// @BasePath /api/v1
// @Summary Create an opening
// @Description Create an opening
// @Tags opening
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Create Opening Request"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %s", err.Error())
		sendErr(ctx, err, http.StatusBadRequest)
		return
	}

	opening := schemas.Job{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %s", err.Error())
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "createOpening", opening)
}
