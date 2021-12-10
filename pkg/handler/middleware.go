package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader     = "Authorization"
	userIdKey               = "userId"
	correctHeaderPartsCount = 2
	tokenIndex              = 1
)

// checkUserIdentity checks whether the user is authorized or not and set user id into context
func (h *Handler) checkUserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != correctHeaderPartsCount {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth token")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[tokenIndex])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	c.Set(userIdKey, userId)
}

// getUserId returns id of authorized user and error
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userIdKey)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is not found")
		return 0, errors.New("user id is not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is not of valid type")
		return 0, errors.New("user id is not of valid type")
	}

	return idInt, nil
}
