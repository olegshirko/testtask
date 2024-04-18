package v1

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
