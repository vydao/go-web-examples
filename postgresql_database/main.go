package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "demo"
	dbname = "demo"
)

func createTableUsers(db *sql.DB) {
	query := `
		CREATE TABLE users (
			id serial PRIMARY KEY,
			username VARCHAR ( 50 ) UNIQUE NOT NULL,
			password VARCHAR ( 50 ) NOT NULL,
			created_at TIMESTAMP
		);`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successful created!")
}

func insertUser(db *sql.DB) {
	username := "Johndoe4"
	password := "secret"
	createdAt := time.Now()
	var id int

	err := db.QueryRow(`INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3) RETURNING id`,
		username, password, createdAt).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Last insert ID: %d\n", id)
}

func getUser(db *sql.DB, uid int64) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)
	err := db.QueryRow(`SELECT * FROM users WHERE id = $1`, uid).Scan(&id, &username, &password, &createdAt)
	if err != nil {
		panic(err)
	}

	fmt.Println(id, username, password, createdAt)
}

func getAllUsers(db *sql.DB) {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", users)
}

func deleteUser(db *sql.DB, id int64) {
	_, err := db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		panic(err)
	}
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successful connected!")

	// Create new table
	// createTableUsers(db)

	// Insert new user
	// insertUser(db)

	// Query a single user
	// getUser(db, 5)

	// Query all users
	// getAllUsers(db)

	// Delete an user
	// deleteUser(db, 6)
}
