package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/meet4041/students-api/internal/config"
)

func main() {
	// load
	cfg := config.MustLoad()

	// database setup

	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Meet Gandhi"))
	})

	// setup server
	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}

	slog.Info("Server Started", slog.String("address: ", cfg.HttpServer.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error %s", err)
		}
	}()
	<-done

	slog.Info("shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown the server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")

}
