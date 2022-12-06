package http

import (
	"encoding/json"
	"net/http"
)

//func (h *Handler) initJobRoutes(api *gin.RouterGroup) {
//	load := api.Group("/job")
//	{
//		load.POST("/", h.loadJob)
//		load.GET("/status", h.getJobStatus)
//	}
//}

func (h *Handler) loadJob(writer http.ResponseWriter, request *http.Request) {
	var uris []string

	if err := json.NewDecoder(request.Body).Decode(&uris); err != nil {
		newErrorResponse(writer, http.StatusBadRequest, "Invalid input body")
		return
	}

	jobStatus, err := h.services.JobService.Create(request.Context(), uris)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(writer).Encode(jobStatus); err != nil {
		return
	}
}

func getJobStatus(writer http.ResponseWriter, request *http.Request) {
}
