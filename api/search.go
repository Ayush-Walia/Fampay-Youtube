package api

import (
	"fmt"
	"net/http"

	"github.com/gookit/slog"
)

func search(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "search API")
	if err != nil {
		slog.Error(err)
	}
}
