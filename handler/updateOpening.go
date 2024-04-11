package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gojobs/schemas"
)

// @BasePath /api/v1
// @Summary Update an opening
// @Description Update an opening
// @Tags opening
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Param request body UpdateOpeningRequest true "Update Opening Request"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validtion error: %s", err.Error())
		sendErr(ctx, err, http.StatusBadRequest)
		return
	}

	id := ctx.Query("id")
	if id == "" {
		logger.Errorf("error: id is required")
		sendErr(ctx, errParamIsRequired("id", "string"), http.StatusBadRequest)
		return
	}

	opening := schemas.Job{}
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("error: %s", err.Error())
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "updateOpening", opening)
}
