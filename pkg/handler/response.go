package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	todo "todolistBackend"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type GetAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

type GetAllItemsResponse struct {
	Data []todo.TodoItem `json:"data"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}
