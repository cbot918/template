package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store DB, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", acc.Number)

	return acc
}

func seedAccounts(s DB) {
	seedAccount(s, "anthony", "GG", "888888")
}

func main() {

	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	cfg, err := GetConfig()
	if err != nil {
		log.Fatal()
	}

	db, err := NewPSQLDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.InitTable(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the database")
		// seed stuff
		seedAccounts(db)
	}

	api := NewAPI(cfg.HOST, db)
	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
