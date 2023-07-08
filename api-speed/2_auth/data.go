package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB interface {
	InitTable() error
	CreateUser(*User) error
	GetUsers() ([]*User, error)
	UpdateUser(*User) (*User, error)
	DeleteUser(int) error
	GetUserByID(int) (*User, error)
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
		CREATE TABLE IF NOT EXISTS users (
			id serial not null primary key,
			name varchar(50) not null,
			email varchar(50) not null,
			password varchar(50) not null,
			created_at timestamp
		)
	`
	_, err := p.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (p *PSQLDB) CreateUser(u *User) error {
	query := `
		INSERT INTO users(
			name, email, password, created_at
		) VALUES (
			$1,$2,$3,$4
		)
	`
	_, err := p.DB.Exec(query, u.Name, u.Email, u.Password, u.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
func (p *PSQLDB) GetUsers() ([]*User, error) {
	query := `
		SELECT * FROM users
	`
	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}
	users := []*User{}
	for rows.Next() {
		user, err := scanIntoUsers(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (p *PSQLDB) UpdateUser(*User) (*User, error) {
	return &User{}, nil
}
func (p *PSQLDB) DeleteUser(id int) error {
	query := `DELETE from users where id = $1`
	_, err := p.DB.Exec(query, id)
	return err
}
func (p *PSQLDB) GetUserByID(id int) (*User, error) {
	query := `select * from users where id=$1`
	rows, err := p.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUsers(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func scanIntoUsers(rows *sql.Rows) (*User, error) {
	user := &User{}
	err := rows.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	return user, err
}
