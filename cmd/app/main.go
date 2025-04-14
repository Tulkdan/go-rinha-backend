package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Tulkdan/go-rinha-backend/src"
	"github.com/Tulkdan/go-rinha-backend/src/db"
	"github.com/jackc/pgx/v5"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func run() error {
	ctx := context.Background()

	connStr := fmt.Sprintf("postgres://%s@%s:%s/%s",
		getEnv("DB_USER", "postgres"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_NAME", "rinha"),
	)
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	if err := src.NewHTTPServer(":8000", ctx, queries).ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
