package main

import (
	"fmt"
	"log/slog"
	"os"
)

var (
	listenAddr = "0:0:0:0"
	listenPort = "8080"
	dbUri      = ""
)

func readEnv() error {
	if addr := os.Getenv("LISTEN_ADDR"); addr != "" {
		listenAddr = addr
	}
	if port := os.Getenv("LISTEN_PORT"); port != "" {
		listenPort = port
	}

	if db := os.Getenv("DB_URI"); db != "" {
		dbUri = db
	} else {
		return fmt.Errorf("DB_URI environment variable not set")
	}

	return nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Starting catalog service")
	err := readEnv()
	if err != nil {
		logger.Error("Failed to read environment variable", "error", err)
		os.Exit(1)
	}

}
