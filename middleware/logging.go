package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Catat waktu mulai
		start := time.Now()

		// Panggil handler berikutnya dalam chain
		next.ServeHTTP(w, r)

		// Catat waktu selesai dan log waktu yang diambil
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
