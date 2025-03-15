package main

import (
	"log"
)

func main() {
	cfg := config{
		addr: ":8080",
		env:  "development",
		db: dbConfig{
			addr:         "postgres://admin:adminpassword@localhost:5432/urls",
			maxOpenConns: 30,
			maxIdleConns: 30,
			maxIdleTime:  "15m",
		},
		apiURL: "http://localhost:8080",
	}

	app := &application{
		config: cfg,
	}

	mux := app.serve()
	log.Fatal(app.run(mux))
}
