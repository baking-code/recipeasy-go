package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/baking-code/recipeasy-go/internal/rest"
	"github.com/baking-code/recipeasy-go/internal/service"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

type Server struct {
	*http.Server
	l *httplog.Logger
}

func MakeServer(f func() int) *Server {
	logger := httplog.NewLogger("average-number-service", httplog.Options{
		LogLevel:         slog.LevelDebug,
		JSON:             true,
		Concise:          true,
		RequestHeaders:   true,
		ResponseHeaders:  true,
		MessageFieldName: "message",
		LevelFieldName:   "level",
		TimeFieldFormat:  time.RFC1123,
		QuietDownPeriod:  10 * time.Second,
	})
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger, []string{"/ping"}))
	r.Use(middleware.Heartbeat("/ping"))

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			httplog.LogEntrySetField(ctx, "user", slog.StringValue("user1"))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	service := service.NewSimpleServiceWithInMemoryDao()
	handler := rest.NewHandler(service)
	handler.Register(r)
	port := ":3333"
	server := http.Server{
		Addr:         "localhost" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Server: &server,
		l:      logger,
	}
}

// Start runs ListenAndServe on the http.Server with graceful shutdown
func (srv *Server) Start() {
	srv.l.Info("Starting server...")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.l.Error("Could not listen on "+srv.Addr, err)
		}
	}()
	srv.l.Info("Server is ready to handle requests " + srv.Addr)
}

func (srv *Server) CloseOnSignal() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	srv.l.Info("Server is shutting down: " + sig.String())

	srv.Close()
}

func (srv *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		srv.l.Error("Could not gracefully shutdown the server")
	}
	srv.l.Info("Server stopped")
}
