package main

import (
	"ing2-tp1/internal"
)

func main() {
	a := internal.App{}
	a.Initialize(
		"postgres",
		"1234",
		"guita")
	a.Run("8000")
}
