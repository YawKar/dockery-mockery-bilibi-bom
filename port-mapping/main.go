package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer cancel()

	handler := http.NewServeMux()

	handler.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello!"))
	})

	server := &http.Server{
		Addr:    ":5050",
		Handler: handler,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		<-ctx.Done()
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		if err := server.Shutdown(ctx); err != nil {
			slog.Error("error trying to shutdown server", slog.String("err", err.Error()))
		}
	}()

	slog.Info("starting to listen", slog.String("addr", server.Addr))
	if err := server.ListenAndServe(); err != nil {
		slog.Error("error listening and serving", slog.String("err", err.Error()))
	}
}
