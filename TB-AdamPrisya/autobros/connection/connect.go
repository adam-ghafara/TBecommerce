package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=12345 dbname=panamora sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Database Connected!")

	return db, nil
}
