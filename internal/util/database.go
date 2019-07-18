package util

import (
	"database/sql"

	_ "github.com/lib/pq" // Driver
)

// Database is the database connection
var Database *sql.DB

// Initialize connects to the database
func Initialize(connectionString string) {
	var err error

	Database, err = sql.Open("postgres", connectionString)

	Check(err)

	pingErr := Database.Ping()

	Check(pingErr)
}
