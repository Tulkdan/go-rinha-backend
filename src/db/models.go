// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Person struct {
	ID        pgtype.UUID
	Name      pgtype.Text
	Nickname  pgtype.Text
	Birthdate pgtype.Timestamp
	Stacks    []string
}
