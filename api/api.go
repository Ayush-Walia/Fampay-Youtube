package api

import "github.com/gorilla/mux"

func InitHandlers(router *mux.Router) {
	router.HandleFunc("/health_check", healthCheck).Methods("GET")

	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/videos", videos).Methods("GET")
	subRouter.HandleFunc("/search", search).Methods("GET")
}
