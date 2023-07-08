package main

import (
	"fmt"
	"log"
)

var lg = fmt.Println

func main() {

	cfg, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	store, err := NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = store.Init()
	if err != nil {
		log.Fatal(err)
	}

	qauth := NewQauth(":3005", store)

	qauth.Run()
}
