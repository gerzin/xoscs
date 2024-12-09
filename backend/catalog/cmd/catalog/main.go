package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Starting catalog service")

	config, err := readEnv()

	if err != nil {
		logger.Error("Failed to read environment variable", "error", err)
		os.Exit(1)
	}

}
