package http

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) loadJob(writer http.ResponseWriter, request *http.Request) {
	var uris []string

	if err := json.NewDecoder(request.Body).Decode(&uris); err != nil {
		newInvalidBodyResponse(writer, err.Error())
		return
	}

	jobStatus, err := h.services.JobService.Create(request.Context(), uris)
	if err != nil {
		newInternalServerErrorResponse(writer, err.Error())
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(writer).Encode(jobStatus); err != nil {
		return
	}
}

func (h *Handler) getJobStatus(writer http.ResponseWriter, request *http.Request) {
}
