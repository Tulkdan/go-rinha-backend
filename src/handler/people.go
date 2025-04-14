package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tulkdan/go-rinha-backend/src/db"
	"github.com/Tulkdan/go-rinha-backend/src/domain"
	"github.com/Tulkdan/go-rinha-backend/src/dto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

var jsonContentType = "application/json"

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
		http.Error(w, domain.ErrIdFailedToParse.Error(), http.StatusBadRequest)
		return
	}

	dbPerson, err := h.db.GetPerson(req.Context(), pgtype.UUID{Bytes: ID, Valid: true})
	if err != nil {
		fmt.Printf("Error getting person %s\n", err)
		http.Error(w, domain.ErrUserNotFound.Error(), http.StatusBadRequest)
		return
	}

	person := domain.NewPersonFromDB(dbPerson)

	w.Header().Set("Content-type", jsonContentType)
	json.NewEncoder(w).Encode(dto.FromPerson(person))
}

func (h *httpServer) HandlePost(w http.ResponseWriter, req *http.Request) {
	var newPersonInput *dto.PersonInput
	err := json.NewDecoder(req.Body).Decode(&newPersonInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	person := dto.ToPerson(newPersonInput)

	_, err = h.db.CreatePerson(req.Context(), person.SaveDB())
	if err != nil {
		fmt.Println(err)
		http.Error(w, domain.ErrInsertPerson.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-type", jsonContentType)
	json.NewEncoder(w).Encode(dto.FromPerson(person))
}
