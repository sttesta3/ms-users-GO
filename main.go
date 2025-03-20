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
		"postgres",
		"1234",
		"ingsoft2")
	a.Run(serverAddr)
}
