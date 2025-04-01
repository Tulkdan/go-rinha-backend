package main

import (
	"context"
	"log"

	"github.com/Tulkdan/go-rinha-backend/src"
	"github.com/Tulkdan/go-rinha-backend/src/db"
	"github.com/jackc/pgx/v5"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres@localhost:5432/rinha")
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
