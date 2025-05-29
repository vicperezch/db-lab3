package main

import (
	"lab3-backend/handlers"
	"lab3-backend/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	utils.ConnectDB()

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.Get("/api/enrollments", handlers.GetEnrollmentsHandler)

	log.Println("Server starting on http://localhost:8080...")
	http.ListenAndServe(":8080", r)
}
