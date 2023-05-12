package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	db_type  = "postgres"
	user     = "postgres"
	password = "12345"
	host     = "localhost"
	port     = "5432"
	db       = "testdb"
)

func main() {

	connStr := fmt.Sprintf("%s://%s:%s@localhost:%s/%s?sslmode=disable", db_type, user, password, port, db)
	// fmt.Println(connStr)
	conn, err := sql.Open(db_type, connStr)
	if err != nil {
		fmt.Println("sql.Open failed")
		panic(err)
	}
	defer conn.Close()
	fmt.Println("postgres conn established")

	err = conn.Ping()
	if err != nil {
		fmt.Println("conn.Ping failed")
		panic(err)
	}
	fmt.Println("ping success")
}
