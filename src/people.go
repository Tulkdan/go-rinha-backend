package src

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Tulkdan/go-rinha-backend/src/db"
	"github.com/jackc/pgx/v5/pgtype"
)

var jsonContentType = "application/json"

var ErrIdNotFound = fmt.Errorf("ID not found")
var ErrInsertPerson = fmt.Errorf("Error inserting person")

type httpServer struct {
	ctx context.Context
	db  *db.Queries
}

func NewPeopleRouter(ctx context.Context, db *db.Queries) *httpServer {
	return &httpServer{
		ctx: ctx,
		db:  db,
	}
}

func (h *httpServer) HandleGet(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")

	person, err := h.db.GetPerson(h.ctx, id)
	if err != nil {
		fmt.Printf("Error getting person %s\n", err)
		http.Error(w, ErrIdNotFound.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", jsonContentType)
	json.NewEncoder(w).Encode(person)
}

type InsertPerson struct {
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Birthdate time.Time `json:"birthdate"`
	Stack     []string  `json:"stack"`
}

type IDDocument struct {
	Id string `json:"id"`
}

func (h *httpServer) HandlePost(w http.ResponseWriter, req *http.Request) {
	var newPerson InsertPerson
	err := json.NewDecoder(req.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(newPerson.Stack)

	person, err := h.db.CreatePerson(h.ctx, db.CreatePersonParams{
		Name:      pgtype.Text{String: newPerson.Name},
		Nickname:  pgtype.Text{String: newPerson.Nickname},
		Birthdate: pgtype.Timestamp{Time: newPerson.Birthdate},
		Stacks:    buf.Bytes(),
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, ErrInsertPerson.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-type", jsonContentType)
	json.NewEncoder(w).Encode(IDDocument{Id: person.ID})
}
