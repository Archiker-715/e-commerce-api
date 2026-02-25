package main

import (
	"log"
	"net/http"

	"github.com/Archiker-715/e-commerce-api/internal/app"
)

func main() {
	app := app.Run()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
