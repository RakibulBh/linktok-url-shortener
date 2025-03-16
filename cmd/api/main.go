package main

import (
	"log"
	"os"

	"github.com/RakibulBh/linktok/internal/db"
	"github.com/RakibulBh/linktok/internal/store"
	_ "github.com/lib/pq"
)

func main() {
	if os.Getenv("ENV") == "development" {
		loadEnv()
	}

	cfg := config{
		addr: ":" + GetEnv("PORT", "8080"),
		env:  GetEnv("ENV", "development"),
		db: dbConfig{
			addr:         GetEnv("DB_ADDR", "postgres://admin:adminpassword@localhost:5432/urls?sslmode=disable"),
			maxOpenConns: 30,
			maxIdleConns: 30,
			maxIdleTime:  "15m",
		},
		apiURL: GetEnv("API_URL", "http://localhost:8080"),
	}

	// log the environment
	log.Printf("Environment: %s", cfg.env)

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
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
