package main

import (
	"ing2-tp1/internal"
	"os"
)

func main() {
	a := internal.App{}
	serverAddr := os.Getenv("SERVER_ADDRESS")
	if serverAddr != "" {	// Solo si pude obtener mi direccion hace algo 
		a.Initialize(
			os.Getenv("MONGO_USER"),
			os.Getenv("MONGO_PASSWORD"),
			os.Getenv("MONGO_DB"),
		)
		a.Run(serverAddr)
	}
}
