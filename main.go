package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HaikalRFadhilahh/auth/controllers/auth"
	"github.com/HaikalRFadhilahh/auth/helper"
	"github.com/HaikalRFadhilahh/auth/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Inject Go .env File
	godotenv.Load()

	// Declation Mux Router
	router := mux.NewRouter()

	// Middleware
	router.Use(middleware.ErrorHandle)

	// Route Path
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://haik.my.id", http.StatusPermanentRedirect)
	})
	router.HandleFunc("/login", auth.Login).Methods("POST")

	// Listen and Serve Router Mux
	PORT := helper.GetEnv("PORT", "8000")
	HOST := helper.GetEnv("HOST", "0.0.0.0")
	listenString := fmt.Sprintf("%s:%s", HOST, PORT)
	err := http.ListenAndServe(listenString, router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
