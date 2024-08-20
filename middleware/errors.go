package middleware

import (
	"encoding/json"
	"net/http"
)

type errorHandle struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func ErrorHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		defer func() {
			ErrorResponse := errorHandle{
				StatusCode: 500,
				Status:     "error",
				Message:    "Internal Server Error",
			}
			err := recover()
			if err != nil {
				json.NewEncoder(w).Encode(ErrorResponse)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}
