package main

import (
	"log"

	"github.com/arunko350/geebee/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
