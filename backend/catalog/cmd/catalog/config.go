package main

// Configuration settings for the catalog service.
type Config struct {
	// Address on which the service will listen.
	ListenAddr string

	// Port on which the service will listen.
	ListenPort string

	// WURI for connecting to the database.
	DBUri string
}
