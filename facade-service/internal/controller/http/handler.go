package http

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/service"
	token_manager "github.com/tuxoo/smart-loader/facade-service/internal/util/token-manager"
	"net/http"
	"time"
)

type Handler struct {
	userService     service.IUserService
	jobService      service.IJobService
	jobStageService service.IJobStageService
	downloadService service.IDownloadService
	tokenManager    token_manager.TokenManager
}

func NewHandler(
	cfg *config.HTTPConfig,
	userService service.IUserService,
	jobService service.IJobService,
	jobStageService service.IJobStageService,
	downloadService service.IDownloadService,
	tokenManager token_manager.TokenManager,
) *gin.Engine {
	handler := &Handler{
		userService:     userService,
		jobService:      jobService,
		jobStageService: jobStageService,
		downloadService: downloadService,
		tokenManager:    tokenManager,
	}

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

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	handler.initMetrics(router)
	handler.initApi(router)

	return router
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	h.initMetrics(router)
	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		h.initUserRoutes(api)
		h.initJobRoutes(api)
	}
}

func (h *Handler) initMetrics(router *gin.Engine) {
	metrics := router.Group("/metrics")
	{
		h.initMetricRoutes(metrics)
	}
}
