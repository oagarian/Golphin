package user

import "github.com/google/uuid"

type User interface {
	ID() *uuid.UUID
	Username() string
	Email() string
	Password() string

	SetID(id *uuid.UUID) error
	SetUsername(username string) error
	SetEmail(email string) error
	SetPassword(password string) error
}