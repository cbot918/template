package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB interface {
	InitTable() error
	CreateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	UpdateAccount(*Account) (*Account, error)
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
	GetAccountByNumber(int) (*Account, error)
}

type PSQLDB struct {
	DB *sql.DB
}

func NewPSQLDB(cfg *Config) (DB, error) {
	db, err := sql.Open("postgres", cfg.PSQL_URL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PSQLDB{
		DB: db,
	}, nil
}

func (p *PSQLDB) InitTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS Accounts (
			id serial not null primary key,
			first_name varchar(100) not null,
			last_name varchar(100) not null,
			number serial ,
			encrypted_password varchar(255),
			balance serial ,
			created_at timestamp
		)
	`
	_, err := p.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (p *PSQLDB) CreateAccount(acc *Account) error {
	query := `
		INSERT INTO Accounts(
			first_name, last_name, number, encrypted_password, balance, created_at
		) VALUES (
			$1,$2,$3,$4,$5,$6
		)
	`
	_, err := p.DB.Exec(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.EncryptedPassword,
		acc.Balance,
		acc.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
func (p *PSQLDB) GetAccounts() ([]*Account, error) {
	query := `
		SELECT * FROM Accounts
	`
	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}
	Accounts := []*Account{}
	for rows.Next() {
		Account, err := scanIntoAccounts(rows)
		if err != nil {
			return nil, err
		}
		Accounts = append(Accounts, Account)
	}
	return Accounts, nil
}
func (p *PSQLDB) UpdateAccount(*Account) (*Account, error) {
	return &Account{}, nil
}
func (p *PSQLDB) DeleteAccount(id int) error {
	query := `DELETE from Accounts where id = $1`
	_, err := p.DB.Exec(query, id)
	return err
}

func (p *PSQLDB) GetAccountByNumber(number int) (*Account, error) {
	query := `select * from Accounts where number=$1`
	rows, err := p.DB.Query(query, number)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccounts(rows)
	}

	return nil, fmt.Errorf("account with number [%d] not found", number)
}

func (p *PSQLDB) GetAccountByID(id int) (*Account, error) {
	query := `select * from Accounts where id=$1`
	rows, err := p.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccounts(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func scanIntoAccounts(rows *sql.Rows) (*Account, error) {
	Account := &Account{}
	err := rows.Scan(
		&Account.ID,
		&Account.FirstName,
		&Account.LastName,
		&Account.Number,
		&Account.EncryptedPassword,
		&Account.Balance,
		&Account.CreatedAt,
	)

	return Account, err
}
