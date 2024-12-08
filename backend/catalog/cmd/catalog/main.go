package main

import (
	"log/slog"
	"os"
)

var (
	listenAddr = "0:0:0:0"
	listenPort = "8080"
)

func readEnv() {
	if addr := os.Getenv("LISTEN_ADDR"); addr != "" {
		listenAddr = addr
	}
	if port := os.Getenv("LISTEN_PORT"); port != "" {
		listenPort = port
	}
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Starting catalog service")
	readEnv()
}
