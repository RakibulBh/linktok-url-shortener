package main

import (
	"context"
	"log"
	"os"

	"github.com/RakibulBh/linktok/internal/db"
	"github.com/RakibulBh/linktok/internal/store"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
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

	// Init redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
		redis:  client,
	}

	mux := app.serve()
	log.Fatal(app.run(mux))
}
