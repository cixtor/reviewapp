package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const approvedReview int = 1
const notApprovedReview int = 0

// Application holds the connection with the database and HTTP handlers.
type Application struct {
	db *sql.DB
}

// Response defines the structure of the JSON output.
type Response struct {
	OK   bool        `json:"ok"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// Reviews is a list of product ratings.
type Reviews []ReviewPublic

// Review defines the data for a product rating.
type Review struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	Approved  bool      `json:"approved"`
	Timestamp time.Time `json:"timestamp"`
}

// ReviewPublic defines the data that can be publicly accessed.
type ReviewPublic struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	ShortDate string `json:"shortdate"`
	Comment   string `json:"comment"`
	Score     int    `json:"score"`
}

// NewApp creates a new instance of the application.
func NewApp(database string) *Application {
	folder := filepath.Dir(os.Args[0]) /* go 1.7.5 */
	conn, err := sql.Open("sqlite3", folder+"/"+database)

	if err != nil {
		log.Fatal("SQLite open; ", err)
		return &Application{}
	}

	log.Println("Database:", folder+"/"+database)

	query := `
	CREATE TABLE IF NOT EXISTS reviews (
		id INTEGER PRIMARY KEY,
		uid TEXT,
		name TEXT,
		email TEXT,
		rating INTEGER,
		comment TEXT,
		approved INTEGER,
		timestamp TIMESTAMP
	);
	`

	if _, err := conn.Exec(query); err != nil {
		log.Fatal("Database setup; ", err)
		return &Application{}
	}

	return &Application{db: conn}
}
