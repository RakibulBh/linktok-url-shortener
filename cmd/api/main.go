package main

import (
	"database/sql"
	"log"

	"github.com/RakibulBh/linktok/internal/store"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config{
		addr: ":8080",
		env:  "development",
		db: dbConfig{
			addr:         "postgres://admin:adminpassword@localhost:5432/urls?sslmode=disable",
			maxOpenConns: 30,
			maxIdleConns: 30,
			maxIdleTime:  "15m",
		},
		apiURL: "http://localhost:8080",
	}

	db, err := sql.Open("postgres", cfg.db.addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.serve()
	log.Fatal(app.run(mux))
}
