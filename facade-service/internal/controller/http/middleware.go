package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	strId, err := h.parseAuthHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Set(userCtx, id)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.tokenManager.ParseToken(headerParts[1])
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user doesn't exist in context")
	}

	return id.(int), nil
}
