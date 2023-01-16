package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

const APPLICATION_ZIP = "application/zip"

func (h *Handler) initJobRoutes(api *gin.RouterGroup) {

	jobs := api.Group("/job", h.userIdentity)
	{
		jobs.POST("/", h.loadJob)
		jobs.GET("/", h.getJobs)
		jobs.GET("/:id/download", h.getDownloads)
	}
}

func (h *Handler) loadJob(c *gin.Context) {
	var urls []string
	if err := c.ShouldBindJSON(&urls); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized user")
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
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized user")
		return
	}

	jobs, err := h.jobService.GetAll(c.Request.Context(), userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func (h *Handler) getDownloads(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized user")
		return
	}

	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "job id was absent")
		return
	}

	jobId, err := uuid.Parse(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "job id was uncorrected")
		return
	}

	content, err := h.downloadService.GetDownloadZip(c.Request.Context(), jobId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "something went wrong")
		return
	}

	c.Data(http.StatusOK, APPLICATION_ZIP, content)
}
