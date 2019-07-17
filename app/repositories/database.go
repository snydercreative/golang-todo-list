package repositories

import (
	"database/sql"
	"golang-todo-list/app/util"

	_ "github.com/lib/pq" // Driver
)

var db *sql.DB

// Initialize connects to the database
func Initialize(connectionString string) {
	var err error

	db, err = sql.Open("postgres", connectionString)

	util.Check(err)
}
