package src

import (
	"context"
	"net/http"

	"github.com/Tulkdan/go-rinha-backend/src/db"
)

func NewHTTPServer(addr string, ctx context.Context, db *db.Queries) *http.Server {
	server := NewPeopleRouter(ctx, db)

	r := &http.ServeMux{}
	r.HandleFunc("GET /pessoas/{id}", server.HandleGet)
	r.HandleFunc("POST /pessoas", server.HandlePost)

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
