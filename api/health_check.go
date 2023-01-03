package api

import (
	"net/http"

	"github.com/Ayush-Walia/Fampay-Youtube/utils"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithString(w, http.StatusOK, "health check is awesome!")
}
