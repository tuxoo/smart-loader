package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initJobRoutes(api *gin.RouterGroup) {

	jobs := api.Group("/job", h.userIdentity)
	{
		jobs.POST("/", h.loadJob)
		jobs.GET("/", h.getJobs)
		jobs.GET("/status", h.getJobStatus)
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

	job, err := h.jobService.Create(c.Request.Context(), userId, urls)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, job)
}

func (h *Handler) getJobs(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized user")
		return
	}

	jobs, err := h.jobService.GetAll(c.Request.Context(), userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func (h *Handler) getJobStatus(c *gin.Context) {

}
