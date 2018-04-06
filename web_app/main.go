package main

import (
	"log"

	"github.com/abdev/huiskluis/web_app/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
