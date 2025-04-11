package main

import (
	"ing2-tp1/internal"
	"os"
)

func main() {
	app := internal.Initialize(
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("MONGO_DB"),
	)
	app.Run(os.Getenv("HOST"), os.Getenv("PORT"))
}
