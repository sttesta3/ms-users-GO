package main

import (
	"ing2-tp1/internal"
	"os"
)

func main() {
	a := internal.App{}
	serverAddr := os.Getenv("SERVER_ADDRESS")
	if serverAddr == "" {
		serverAddr = "localhost:8000"
	}

	a.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PWD"),
		os.Getenv("POSTGRES_DB"),
	)
	a.Run(serverAddr)
}
