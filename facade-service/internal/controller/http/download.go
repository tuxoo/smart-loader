package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) initDownloadRoutes(api *gin.RouterGroup) {

	load := api.Group("/download", h.userIdentity)
	{
		load.GET("/", h.getDownloads)
	}
}

func (h *Handler) getDownloads(c *gin.Context) {
	jobId := c.Query("jobId")
	if jobId == "" {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprint("empty field [jobId]"))
		return
	}

	_, err := uuid.Parse(jobId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprint("incorrect field [jobId]"))
		return
	}
}
