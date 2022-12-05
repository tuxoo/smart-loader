package http

import (
	"github.com/gorilla/mux"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/service"
	"net/http"
)

type Handler struct {
	jobService service.IJobService
}

func NewHandler(jobService service.IJobService) *Handler {
	return &Handler{
		jobService: jobService,
	}
}

func (h *Handler) Init(cfg config.HTTPConfig) *mux.Router {
	router := mux.NewRouter()

	router.Use(
		mux.CORSMethodMiddleware(router),
	)

	router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write([]byte("pong")); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	h.initApi(router)
	//h.initMetrics(router)

	return router
}

func (h *Handler) initApi(router *mux.Router) {
	api := router.PathPrefix("/api/v1/job").Subrouter()

	api.HandleFunc("", h.loadJob).Methods(http.MethodPost)
	api.HandleFunc("/status", getJobStatus).Methods(http.MethodGet)
}

//func (h *Handler) initMetrics(router *gin.Engine) {
//	metrics := router.Group("/metrics")
//	{
//		h.initMetricRoutes(metrics)
//	}
//}
