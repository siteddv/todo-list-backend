package handler

import (
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
