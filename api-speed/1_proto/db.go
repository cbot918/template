package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB interface {
	CreateUser(*User) error
	ReadUser() (*[]User, error)
	UpdateUser(int) error
	DeleteUser(*User) error
	GetUserByID(int) (*User, error)
}

type PostgresDB struct {
	DB *sql.DB
}

// docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
func NewPostgresDB(cfg *Config) (*PostgresDB, error) {
	db, err := sql.Open("postgres", cfg.PSQL_URL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresDB{
		DB: db,
	}, nil
}

func (p *PostgresDB) Init() error {
	query := `
		create table if not exists users(
			id serial primary key, 
			name varchar(50) not null,
			email varchar(50) not null,
			password varchar(50) not null,
			created_at timestamp
		)
	`
	_, err := p.DB.Exec(query)
	return err
}

func (p *PostgresDB) CreateUser(u *User) error {
	query := `
	insert into users(
		name,
		email,
		password,
		created_at
	) values ($1,$2,$3,$4)
	`
	_, err := p.DB.Exec(query, u.Name, u.Email, u.Password, u.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (p *PostgresDB) ReadUser() ([]User, error) {
	query := `
		select * from users
	`
	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}
	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
func (p *PostgresDB) UpdateUser(u *User) error {
	return nil
}
func (p *PostgresDB) DeleteUser(id int) error {
	return nil
}
func (p *PostgresDB) GetUserByID(id int) (*User, error) {
	return &User{}, nil
}
