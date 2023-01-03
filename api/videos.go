package api

import (
	"fmt"
	"net/http"

	"github.com/gookit/slog"
)

func videos(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "videos API")
	if err != nil {
		slog.Error(err)
	}
}
