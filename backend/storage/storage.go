package storage

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/lib/pq" // postgres driver
)

const secretFile = `/run/secrets/db-password`
const queryDropTable = `DROP TABLE IF EXISTS users;`
const queryCreateTable = `
CREATE TABLE IF NOT EXISTS users (
	user_id   SERIAL,
	username  TEXT,
	data      TEXT
);`

// User represents a user an also their web homepage data
type User struct {
	Id   uint
	Name string
	Data string
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

	if _, err := s.db.Exec(queryDropTable); err != nil {
		log.Fatal("Database Table Drop Failed: ", err)
	}

	if _, err := s.db.Exec(queryCreateTable); err != nil {
		log.Fatal("Database Table Creation Failed: ", err)
	}

}

// PrintWholeTable returns a string showing ALL data from the 'users' table
func (s *Storage) PrintWholeTable() string {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		log.Println("db query error: ", err)
		return "ERROR: the database was given a bad query"
	}
	var output string
	for rows.Next() {
		var id int
		var username string
		var data string
		err := rows.Scan(&id, &username, &data)
		if err != nil {
			log.Println("error in PrintWholeTable:", err)
		}
		output += fmt.Sprintln("(", id, username, data, ")")
	}
	return output
}

// AddUser creates and inserts a new user row into the database
func (s *Storage) AddUser(user User) {
	const queryAddUser = `
		INSERT INTO users (user_id, username, data)
		VALUES ($1, $2, $3);`
	s.db.Exec(queryAddUser, user.Id, user.Name, user.Data)
}

// GetUserDataByName fetches a user object from the data base, returns "false"
// if that username does not exist. Assumes that usernames are unique, otherwise
// it will return the first user found in the database.
func (s *Storage) GetUserDataByName(name string) (User, bool) {
	u := User{}
	const queryGetUser = `SELECT * FROM users WHERE username=$1`
	err := s.db.QueryRow(queryGetUser, name).Scan(&u.Id, &u.Name, &u.Data)
	return u, err == nil
}

func connect() *sql.DB {
	bin, err := ioutil.ReadFile(secretFile)
	if err != nil {
		log.Fatal("Unable to read database password file!", err)
	}
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://postgres:%s@db:5432/example?sslmode=disable", string(bin)))
	if err != nil {
		log.Fatal("Unable to open database, check driver name and data source name", err)
	}
	return db
}
