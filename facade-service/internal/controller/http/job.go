package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initJobRoutes(api *gin.RouterGroup) {

	load := api.Group("/job", h.userIdentity)
	{
		load.POST("/", h.loadJob)
		load.GET("/status", h.getJobStatus)
	}
}

func (h *Handler) loadJob(c *gin.Context) {
	var urls []string
	if err := c.ShouldBindJSON(&urls); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized user")
		return
	}

	jobStatus, err := h.jobService.Create(c.Request.Context(), userId, urls)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, jobStatus)
}

func (h *Handler) getJobStatus(c *gin.Context) {

}
