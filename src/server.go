package src

import (
	"context"
	"net/http"

	"github.com/Tulkdan/go-rinha-backend/src/db"
	"github.com/Tulkdan/go-rinha-backend/src/handler"
)

func NewHTTPServer(addr string, ctx context.Context, db *db.Queries) *http.Server {
	server := handler.NewPeopleRouter(db)

	r := &http.ServeMux{}
	r.HandleFunc("GET /pessoas/{id}", server.HandleGet)
	r.HandleFunc("POST /pessoas", server.HandlePost)
	r.HandleFunc("GET /pessoas", server.HandleSearch)

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
