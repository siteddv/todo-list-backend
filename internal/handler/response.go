package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"todolistBackend/internal/model"
)

// errorResponse contains message whether server occurs an error and used for reporting on this one in response on request
type errorResponse struct {
	Message string `json:"message"`
}

// statusResponse contains info about response success on request
type statusResponse struct {
	Status string `json:"status"`
}

// GetAllListsResponse contains data about List of todoItems and used for response on GET request
type GetAllListsResponse struct {
	Data []model.TodoList `json:"data"`
}

// GetAllItemsResponse contains data about todoItems and used for response on GET request
type GetAllItemsResponse struct {
	Data []model.TodoItem `json:"data"`
}

// newErrorResponse is used for creating for notifying whether server occurs error on request
func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}
