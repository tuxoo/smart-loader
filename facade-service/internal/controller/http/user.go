package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
	"net/http"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.POST("/sign-in/", h.signIn)

		authenticated := user.Group("/", h.userIdentity)
		{
			authenticated.GET("/", h.getUserProfile)
			authenticated.GET("/refresh/:token", h.refreshToken)
		}
	}
}

func (h *Handler) signIn(c *gin.Context) {
	var signInDto model.SignInDTO

	if err := c.ShouldBindJSON(&signInDto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.userService.SignIn(c.Request.Context(), signInDto)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (h *Handler) getUserProfile(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized user")
		return
	}

	user, err := h.userService.GetById(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized user")
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) refreshToken(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized user")
		return
	}

	token := c.Param("token")
	if token == "" {
		newErrorResponse(c, http.StatusBadRequest, "refresh token was absent")
		return
	}

	refreshToken, err := uuid.Parse(token)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "refresh token was uncorrected")
		return
	}

	tokens, err := h.userService.RefreshToken(c.Request.Context(), userId, refreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, tokens)
}
