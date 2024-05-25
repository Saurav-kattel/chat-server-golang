package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(cnStr string) (*sqlx.DB, error) {
	// Define connection parameters
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "postgres", "saurav", "x-clone")

	// Open a connection to the database
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
