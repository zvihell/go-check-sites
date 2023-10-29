package pkg

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=88888888 dbname=sites sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, err
}
