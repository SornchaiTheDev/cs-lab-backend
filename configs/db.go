package configs

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(config *Config) *sqlx.DB {

	db, err := sqlx.Open("postgres", config.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
