package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/Ayush-Walia/Fampay-Youtube/api"
	"github.com/Ayush-Walia/Fampay-Youtube/config"
	"github.com/Ayush-Walia/Fampay-Youtube/storage"
	"github.com/gookit/slog"
	"github.com/gorilla/mux"
)

// Server provides an http.Server.
type Server struct {
	*http.Server
}

// Serve loads the application's config file and starts the server
func Serve() {
	conf := config.LoadConfig()
	router := mux.NewRouter()

	api.InitHandlers(router)
	storage.Init(conf)
	Start(conf, router)
}

func Start(conf *config.AppConfig, router *mux.Router) {
	addr := ":" + conf.ServerPort

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		// always returns error. ErrServerClosed on graceful close
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			slog.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	slog.Infof("Listening on %s\n", addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	slog.Info("Shutting down server... Reason:", sig)

	if err := server.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}
