package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gojobs/schemas"
)

// @BasePath /api/v1
// @Summary List openings
// @Description List all job openings
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Job{}

	if err := db.Find(&openings).Error; err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "listOpenings", openings)
}
