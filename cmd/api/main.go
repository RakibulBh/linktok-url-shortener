package main

import (
	"log"
)

func main() {
	cfg := config{
		addr:   ":8080",
		env:    "development",
		apiURL: "http://localhost:8080",
	}

	app := &application{
		config: cfg,
	}

	mux := app.serve()
	log.Fatal(app.run(mux))
}
