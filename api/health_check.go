package api

import (
	"fmt"
	"net/http"

	"github.com/gookit/slog"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "health check is awesome!")
	if err != nil {
		slog.Error(err)
	}
}
