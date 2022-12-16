package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
	"net/http"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.POST("/sign-in/", h.signIn)

		//authenticated := user.Group("/", h.userIdentity)
		//{
		//	authenticated.GET("/profile/", h.getUserProfile)
		//}
	}
}

func (h *Handler) signIn(c *gin.Context) {
	var signInDto model.SignInDTO

	if err := c.ShouldBindJSON(&signInDto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.userService.SignIn(c.Request.Context(), signInDto)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}
