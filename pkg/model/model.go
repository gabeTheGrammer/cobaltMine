package model

import (
	"errors"
	"time"
)

var (
	ErrNoRec = errors.New("models: No isntance of the record")

	ErrInvalidCreds = errors.New("models: invalid credentials")

	ErrDuplicateEmail = errors.New("models: duplicate email")
)

type User struct {
	ID           int
	Name         string
	HashPassword []byte
	timeCreated  time.Time
	Email        string
}
