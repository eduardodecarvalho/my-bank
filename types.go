package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
  FromAccount int `json:"toAccount"`  
  ToAccount   int `json:"fromAccount"`
  Amount      int `json:"amount"`
}

type CreateAccountRequest struct {
  FirstName string  `json:"firstName"`
  LastName  string  `json:"lastName"`
}

type Account struct {
  ID        int      `json:"id"`
  FirstName string    `json:"firstName"`
  LastName  string    `json:"lastName"`
  Number    int       `json:"number"`
  Balance   int       `json:"balance"`
  CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Intn(100000),
    CreatedAt: time.Now().UTC(),
	}
}
