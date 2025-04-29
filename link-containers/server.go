package main

import (
	"context"
	"flag"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	port := flag.Arg(0)
	if port == "" {
		slog.Error("no port was specified!")
		os.Exit(1)
	}

	handler := http.NewServeMux()
	handler.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("got request")
		w.Write([]byte("Hello!"))
	})

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			slog.Error("error trying to shutdown server!", slog.String("error", err.Error()))
		}
	}()

	slog.Info("listening", slog.String("port", port))

	if err := server.ListenAndServe(); err != nil {
		slog.Error("error listening and serving!", slog.String("error", err.Error()))
	}
}
