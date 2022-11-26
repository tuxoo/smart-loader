package http

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) initLoadRoutes(api *gin.RouterGroup) {
	load := api.Group("/load")
	{
		load.POST("/job", h.loadJob)
		load.GET("/status", h.loadStatus)
	}
}

func (h *Handler) loadJob(c *gin.Context) {

}

func (h *Handler) loadStatus(c *gin.Context) {

}
