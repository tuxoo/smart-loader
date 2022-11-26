package http

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/service"
	"net/http"
	"time"
)

type Handler struct {
	jobService service.IJobService
}

func NewHandler(jobService service.IJobService) *Handler {
	return &Handler{
		jobService: jobService,
	}
}

func (h *Handler) Init(cfg config.HTTPConfig) *gin.Engine {
	router := gin.New()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.New(corsConfig),
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	h.initMetrics(router)
	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		h.initJobRoutes(api)
	}
}

func (h *Handler) initMetrics(router *gin.Engine) {
	metrics := router.Group("/metrics")
	{
		h.initMetricRoutes(metrics)
	}
}
