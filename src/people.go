package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Tulkdan/go-rinha-backend/src/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

var jsonContentType = "application/json"

var ErrIdFailedToParse = fmt.Errorf("ID failed to parse")
var ErrUserNotFound = fmt.Errorf("User not found")
var ErrInsertPerson = fmt.Errorf("Error inserting person")

type httpServer struct {
	db *db.Queries
}

func NewPeopleRouter(db *db.Queries) *httpServer {
	return &httpServer{db: db}
}

func (h *httpServer) HandleGet(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	ID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, ErrIdFailedToParse.Error(), http.StatusBadRequest)
		return
	}

	person, err := h.db.GetPerson(req.Context(), pgtype.UUID{Bytes: ID, Valid: true})
	if err != nil {
		fmt.Printf("Error getting person %s\n", err)
		http.Error(w, ErrUserNotFound.Error(), http.StatusBadRequest)
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
	id, err := uuid.NewV7()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	person, err := h.db.CreatePerson(req.Context(), db.CreatePersonParams{
		ID:        pgtype.UUID{Bytes: id, Valid: true},
		Name:      pgtype.Text{String: newPerson.Name, Valid: true},
		Nickname:  pgtype.Text{String: newPerson.Nickname, Valid: true},
		Birthdate: pgtype.Timestamp{Time: newPerson.Birthdate, Valid: true},
		Stacks:    newPerson.Stack,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, ErrInsertPerson.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-type", jsonContentType)
	json.NewEncoder(w).Encode(IDDocument{Id: person.ID.String()})
}
