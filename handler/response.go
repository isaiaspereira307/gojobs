package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gojobs/schemas"
)

func sendErr(ctx *gin.Context, err error, status int) {
	logger.Errorf("error: %s", err.Error())
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{"error": err.Error()})
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"operation": operation, "data": data})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
}

type CreateOpeningResponse struct {
	Message string              `json:"message"`
	Data    schemas.JobResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string              `json:"message"`
	Data    schemas.JobResponse `json:"data"`
}

type ShowOpeningResponse struct {
	Message string              `json:"message"`
	Data    schemas.JobResponse `json:"data"`
}

type UpdateOpeningResponse struct {
	Message string              `json:"message"`
	Data    schemas.JobResponse `json:"data"`
}

type ListOpeningsResponse struct {
	Message string                `json:"message"`
	Data    []schemas.JobResponse `json:"data"`
}
