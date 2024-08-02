package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionPool(postgresDSN string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", postgresDSN)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Connected to Postgres")
	return db
}
