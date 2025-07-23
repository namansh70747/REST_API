package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/namsh70747/Rest_API/internal/config"
	"github.com/namsh70747/Rest_API/internal/http/handlers/student"
	"github.com/namsh70747/Rest_API/internal/sqlite"
)

// load config
func main() {
	cfg := config.MustLoad()

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	_ = storage
	slog.Info("storage initialized", slog.String("path", cfg.Env), slog.String("version", "1.0.0"))
	router := http.NewServeMux()

	router.HandleFunc("/api/students", student.New(storage))
	router.HandleFunc("/api/students/list", student.GetList(storage))
	router.HandleFunc("/api/students/{id}", student.GetById(storage))
	//database setup
	//setup router

	//setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr, // Use the address from config
		Handler: router,
	}

	fmt.Printf("server started %s", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done
	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))

	}
}
