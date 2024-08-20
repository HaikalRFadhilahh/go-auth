package helper

import (
	"encoding/json"
	"net/http"
)

type responseStructure struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}

func ErrorResponse(w http.ResponseWriter, statusCode int, status string, message string, data any) {
	w.Header().Set("Content-Type", "application/json")

	res := responseStructure{
		StatusCode: statusCode,
		Status:     status,
		Message:    message,
		Data:       data,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}
