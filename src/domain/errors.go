package domain

import "fmt"

var (
	ErrIdFailedToParse = fmt.Errorf("ID failed to parse")
	ErrUserNotFound    = fmt.Errorf("User not found")
	ErrInsertPerson    = fmt.Errorf("Error inserting person")
)
