package utils

import (
	"database/sql"
)

func establishConnection() *sql.DB {
	// Establish database connection
	return &sql.DB{}
}