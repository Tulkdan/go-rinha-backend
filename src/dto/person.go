package dto

import (
	"time"

	"github.com/Tulkdan/go-rinha-backend/src/domain"
	"github.com/google/uuid"
)

type PersonInput struct {
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Birthdate time.Time `json:"birthdate"`
	Stack     []string  `json:"stack"`
}

type PersonOutput struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Birthdate time.Time `json:"birthdate"`
	Stack     []string  `json:"stack"`
}

func ToPerson(input *PersonInput) *domain.Person {
	return domain.NewPerson(input.Name, input.Nickname, input.Birthdate, input.Stack)
}

func FromPerson(input *domain.Person) *PersonOutput {
	return &PersonOutput{
		Id:        input.Id,
		Name:      input.Name,
		Nickname:  input.Nickname,
		Birthdate: input.Birthdate,
		Stack:     input.Stack,
	}
}

type PeopleCountOutput struct {
	Dount int64 `json:"count"`
}

func FromPeopleCount(input int64) *PeopleCountOutput {
	return &PeopleCountOutput{Dount: input}
}
