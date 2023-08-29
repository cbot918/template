package internal

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB(cfg *Config) (*sqlx.DB, error) {
	dsn := cfg.DB_URL
	db, err := sqlx.Open(cfg.DB_TYPE, dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
