package main

import (
	"fmt"
	"log/slog"
	"os"
)

func readEnv() (*Config, error) {
	config := &Config{
		ListenAddr: "0:0:0:0",
		ListenPort: "8080",
	}

	if addr := os.Getenv("LISTEN_ADDR"); addr != "" {
		config.ListenAddr = addr
	}
	if port := os.Getenv("LISTEN_PORT"); port != "" {
		config.ListenPort = port
	}

	if db := os.Getenv("DB_URI"); db != "" {
		config.DBUri = db
	} else {
		return nil, fmt.Errorf("DB_URI environment variable not set")
	}

	return config, nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Starting catalog service")

	config, err := readEnv()

	if err != nil {
		logger.Error("Failed to read environment variable", "error", err)
		os.Exit(1)
	}

}
