package reqres

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type LoginSuccessResponse struct {
	Token string `json:"token"`
}

func WriteJSONErrorResponse(w http.ResponseWriter, errorMessage string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(ErrorResponse{errorMessage})
}
