package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gojobs/schemas"
)

// @BasePath /api/v1
// @Summary Delete an opening
// @Description Delete an opening
// @Tags opening
// @Accept json
// @Produce json
// @Param id query string true "Delete Opening Request"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendErr(ctx, errParamIsRequired("id", "string"), http.StatusBadRequest)
		return
	}
	opening := schemas.Job{}
	// Find the opening by id
	if err := db.First(&opening, id).Error; err != nil {
		sendErr(ctx, err, http.StatusNotFound)
		return
	}
	// Delete the opening
	if err := db.Delete(&opening).Error; err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "deleteOpening", opening)
}
