package storage

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const queryInit = `
DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users {
    user_id   SERIAL,
    username  TEXT,
    data      TEXT,
};
`

const secretFile = `/run/secrets/db-password`

// User represents a user an also their web homepage data
type User struct {
	id   int
	name string
	data string
}

// Storage is an abstraction of a database for our application, so that we
// don't have to use sql commands every time we want to use it.
type Storage struct {
	db *sql.DB
}

// Init will open, connect to, and set up the database so that we can begin
// making use of storage. Will panic if there are any problems. Make sure to
// call this before using the Storage functions
func (s *Storage) Init() {
	log.Println("Initializing database...")

	log.Println("Connecting to database...")
	s.db = connect()

	// Is pinging the server for a full minute really necessary?
	for i := 0; i < 60; i++ {
		if err := s.db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if _, err := s.db.Exec(queryInit); err != nil {
		log.Fatal(err)
	}

	// for i := 0; i < 5; i++ {
	// 	if _, err := db.Exec("INSERT INTO blog (title) VALUES ($1);", fmt.Sprintf("Blog post #%d", i)); err != nil {
	// 		log.
	// 	}
	// }

}

func connect() *sql.DB {
	bin, err := ioutil.ReadFile(secretFile)
	if err != nil {
		log.Fatal("Unable to read database password file!")
	}
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://postgres:%s@db:5432/example?sslmode=disable", string(bin)))
	if err != nil {
		log.Fatal("Unable to open database, check driver name and data source name")
	}
	return db
}
