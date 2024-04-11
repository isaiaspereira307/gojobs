package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gojobs/schemas"
)

// @BasePath /api/v1
// @Summary Show opening
// @Description Show an opening
// @Tags opening
// @Accept json
// @Produce json
// @Param id query string true "Show Opening Request"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
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

	sendSuccess(ctx, "showOpening", opening)
}
