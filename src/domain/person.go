package domain

import (
	"time"

	"github.com/Tulkdan/go-rinha-backend/src/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Person struct {
	Id        uuid.UUID
	Name      string
	Nickname  string
	Birthdate time.Time
	Stack     []string
}

func NewPerson(name, nickname string, birthdate time.Time, stack []string) *Person {
	id, _ := uuid.NewV7()

	return &Person{
		Id:        id,
		Name:      name,
		Nickname:  nickname,
		Birthdate: birthdate,
		Stack:     stack,
	}
}

func NewPersonFromDB(fromDb db.Person) *Person {
	return &Person{
		Id:        fromDb.ID.Bytes,
		Name:      fromDb.Name.String,
		Nickname:  fromDb.Nickname.String,
		Birthdate: fromDb.Birthdate.Time,
		Stack:     fromDb.Stacks,
	}
}

func (p *Person) SaveDB() db.CreatePersonParams {
	return db.CreatePersonParams{
		ID:        pgtype.UUID{Bytes: p.Id, Valid: true},
		Name:      pgtype.Text{String: p.Name, Valid: true},
		Nickname:  pgtype.Text{String: p.Nickname, Valid: true},
		Birthdate: pgtype.Timestamp{Time: p.Birthdate, Valid: true},
		Stacks:    p.Stack,
	}
}
