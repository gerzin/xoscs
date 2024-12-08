package main

import (
	"log/slog"
	"os"
)

var (
	listenAddr = "0:0:0:0"
	listenPort = "8080"
)

// readEnv reads the environment variables LISTEN_ADDR and LISTEN_PORT.
// If these environment variables are set, it updates the listenAddr and listenPort
// variables with their respective values. If the environment variables are not set,
// the default values for listenAddr ("0:0:0:0") and listenPort ("8080") are used.
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
