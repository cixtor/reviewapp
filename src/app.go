package main

import (
	"database/sql"
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
