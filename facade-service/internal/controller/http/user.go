package http

import (
	"encoding/json"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
	"net/http"
)

func (h *Handler) signInUser(writer http.ResponseWriter, request *http.Request) {
	var signInDto model.SignInDTO

	if err := json.NewDecoder(request.Body).Decode(&signInDto); err != nil {
		newInvalidBodyResponse(writer, err.Error())
		return
	}

	token, err := h.services.UserService.SignIn(request.Context(), signInDto)
	if err != nil {
		newForbiddenResponse(writer)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(writer).Encode(map[string]any{
		"token": token,
	}); err != nil {
		return
	}
}
