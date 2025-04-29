package main

import (
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	host := flag.Arg(0)
	if host == "" {
		slog.Error("no host specified!")
		os.Exit(1)
	}
	port := flag.Arg(1)
	if port == "" {
		slog.Error("no port specified!")
		os.Exit(1)
	}

	res, err := http.Get(fmt.Sprintf("http://%s:%s/ping", host, port))
	if err != nil {
		slog.Error("request failed!", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		slog.Error("request failed!", slog.Int("statusCode", res.StatusCode))
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("couldn't read all body!", slog.String("error", err.Error()))
		os.Exit(1)
	}

	fmt.Printf("Response: %s\n", string(body))

	if err := res.Body.Close(); err != nil {
		slog.Error("error trying to close request body!", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
