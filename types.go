package main

import (
	"math/rand"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Number    int
	Balance   int
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Intn(100000),
		Balance:   0,
	}
}
