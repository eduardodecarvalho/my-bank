package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
  FirstName string  `json:"firstName"`
  LastName  string  `json:"lastName"`
}

type Account struct {
  ID        uuid.UUID `json:"id"`
  FirstName string    `json:"firstName"`
  LastName  string    `json:"lastName"`
  Number    int       `json:"number"`
  Balance   int       `json:"balance"`
  CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Intn(100000),
    CreatedAt: time.Now().UTC(),
	}
}
