package main

import "log"

func main() {
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

	api := NewAPI(cfg.HOST, db)
	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
