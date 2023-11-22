package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
  GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `create table if not exists ACCOUNT (
    id serial PRIMARY KEY,
    first_name varchar(50),
    last_name varchar(50),
    number serial,
    balance serial,
    created_at timestamp
  );`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `insert into ACCOUNT
  (first_name, last_name, balance, number, created_at)
  values
  ($1, $2, $3, $4, $5);
  `

  resp, err := s.db.Query(query, 
    acc.FirstName, 
    acc.LastName, 
    acc.Balance, 
    acc.Number, 
    acc.CreatedAt,
  )

	fmt.Println(resp)

	return err
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {  
  fmt.Println("Delete Account: ")
  _, err := s.db.Query(`delete from ACCOUNT where id = $1`, id)
  return err
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
  rows, err := s.db.Query(`select * from ACCOUNT where id = $1`, id)

  if err != nil {
    return nil, err
  }

  for rows.Next() {
    return scanIntoAccount(rows)

  }

  return nil, fmt.Errorf("Account %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
  rows, err := s.db.Query("select * from account")

  if err != nil {
    return nil, err
  }

  accounts := []*Account{}
  for rows.Next() {
    account, err := scanIntoAccount(rows)
    if err != nil {
      return nil, err
    }
    accounts = append(accounts, account)
  }

  return accounts, nil

}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
  account := new(Account)
  err := rows.Scan(
    &account.ID, 
    &account.FirstName, 
    &account.LastName, 
    &account.Number,
    &account.Balance, 
    &account.CreatedAt);

  return account, err
}
