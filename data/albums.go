package data

/*
This page is meant to hold connection strings and a few key tables
In the case that this is to actually be hosted we should use the os
package to pull Connection constants from ENV variables.
This will allow this script to be fully controlled and updated
from the Docker Compose yaml or any other orchestration tool that
may be used.
*/
import (
	"database/sql"
	"fmt"

	// "os"
	_ "github.com/lib/pq"
)

/* --------------------------- Database Constants --------------------------- */
// Declare Database information -> Update this to reflect your database
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// This is the commented out os ENV version in case there is a need to quickly
// Update
// const (
// 	host	 = os.Getenv("DB_HOST")
// 	port	 = os.Getenv("DB_PORT")
// 	user     = os.Getenv("DB_USER")
// 	password = os.Getenv("DB_PASSWD")
// 	dbname   = os.Getenv("DB_NAME")
// )

// Generate Connection string from Database constants
var psqlInfo string = fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

// Open a connection to the Database
func OpenConn() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

/* -------------------------- Table Representations ------------------------- */
// Album Table
type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Create an insert type that does not require id (Since SERIAL type in
// postgres will allocate the id dynamically)
type AlbumInsert struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
