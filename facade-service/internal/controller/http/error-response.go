package http

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type errorResponse struct {
	ErrorTime string `json:"errorTime" example:"2022-06-07 22:22:20"`
	Message   string `json:"message" example:"Token is expired"`
}

func newErrorResponse(writer http.ResponseWriter, statusCode int, message string) {
	logrus.Error(message)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	if err := json.NewEncoder(writer).Encode(errorResponse{
		ErrorTime: time.Now().Format(timeFormat),
		Message:   message,
	}); err != nil {
		return
	}
}

func newInternalServerErrorResponse(writer http.ResponseWriter, message string) {
	newErrorResponse(writer, http.StatusInternalServerError, fmt.Sprintf("Internal server error [%s]", message))
}

func newInvalidBodyResponse(writer http.ResponseWriter, message string) {
	newErrorResponse(writer, http.StatusBadRequest, fmt.Sprintf("Invalid input body [%s]", message))
}

func newEntityNotFoundResponse(writer http.ResponseWriter, message string) {
	newErrorResponse(writer, http.StatusNotFound, fmt.Sprintf("Entity not found [%s]", message))
}

func newForbiddenResponse(writer http.ResponseWriter) {
	newErrorResponse(writer, http.StatusForbidden, "Access denied")
}
