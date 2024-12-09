package main

import (
	"fmt"
	"os"
)

// Configuration settings for the catalog service.
type Config struct {
	// Address on which the service will listen.
	ListenAddr string

	// Port on which the service will listen.
	ListenPort string

	// URI for connecting to the database.
	DBUri string
}

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
