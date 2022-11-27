package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initJobRoutes(api *gin.RouterGroup) {
	load := api.Group("/job")
	{
		load.POST("/", h.loadJob)
		load.GET("/status", h.getJobStatus)
	}
}

func (h *Handler) loadJob(c *gin.Context) {
	var uris []string
	if err := c.ShouldBindJSON(&uris); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	jobStatus, err := h.jobService.Create(context.TODO(), uris)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, jobStatus)
}

func (h *Handler) getJobStatus(c *gin.Context) {

}
