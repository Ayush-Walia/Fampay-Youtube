package api

import (
	"github.com/Ayush-Walia/Fampay-Youtube/middleware"
	"github.com/gorilla/mux"
)

// InitHandlers initializes the HTTP handlers for the API.
func InitHandlers(router *mux.Router) {
	router.HandleFunc("/health_check", healthCheck).Methods("GET")

	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.Use(middleware.ContextMiddleware)
	subRouter.HandleFunc("/videos", videos).Methods("GET")
	subRouter.HandleFunc("/search", search).Methods("GET")
}
