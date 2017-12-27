package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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

// Index handles the HTTP requests to the root of the website.
func (app *Application) Index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Time int64
	}{
		Time: time.Now().Unix(),
	}

	t, err := template.ParseFiles("views/index.tpl")

	if err != nil {
		log.Println("Template.Parse:", err)
		http.Error(w, "Internal Server Error 0x0178", http.StatusInternalServerError)
	}

	if err := t.Execute(w, data); err != nil {
		log.Println("Template.Execute:", err)
		http.Error(w, "Internal Server Error 0x0183", http.StatusInternalServerError)
	}
}

// Admin handles the HTTP requests to the admin interface.
func (app *Application) Admin(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Time int64
	}{
		Time: time.Now().Unix(),
	}

	t, err := template.ParseFiles("views/index.tpl")

	if err != nil {
		log.Println("Template.Parse:", err)
		http.Error(w, "Internal Server Error 0x0178", http.StatusInternalServerError)
	}

	if err := t.Execute(w, data); err != nil {
		log.Println("Template.Execute:", err)
		http.Error(w, "Internal Server Error 0x0183", http.StatusInternalServerError)
	}
}

// ReviewsList responds with a list of reviews for a specific product.
func (app *Application) ReviewsList(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("uid")
	query := "SELECT * FROM reviews WHERE UID = ? AND approved = ? ORDER BY timestamp DESC"
	rows, err := app.db.Query(query, uid, approvedReview)

	if err != nil {
		log.Println("reviews list;", err)
		http.Error(w, "cannot retrieve reviews", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var item Review
	var list Reviews

	for rows.Next() {
		item = Review{}

		if err := rows.Scan(
			&item.ID,
			&item.UID,
			&item.Name,
			&item.Email,
			&item.Rating,
			&item.Comment,
			&item.Approved,
			&item.Timestamp); err != nil {
			log.Println("read review;", err)
			continue
		}

		list = append(list, ReviewPublic{
			ID:        item.ID,
			Name:      item.Name,
			Avatar:    app.Gravatar(item.Email),
			ShortDate: item.Timestamp.Format(`Jan 02, 2006`),
			Score:     item.Rating,
			Comment:   item.Comment,
		})
	}

	json.NewEncoder(w).Encode(Response{OK: true, Data: list})
}

// ReviewsSave responds with a list of reviews for a specific product.
func (app *Application) ReviewsSave(w http.ResponseWriter, r *http.Request) {
	stmt, err := app.db.Prepare(
		`INSERT INTO reviews (uid, name, email, rating, comment,
		approved, timestamp) VALUES (?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		log.Println("review insert;", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	var rating int
	ratingStr := r.Form.Get("rating")
	ratingNum, err := strconv.Atoi(ratingStr)
	if err != nil {
		log.Println("invalid rating;", ratingStr, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ratingNum < 0 {
		rating = 0
	} else if ratingNum > 10 {
		rating = 10
	} else {
		rating = ratingNum
	}

	if _, err := stmt.Exec(
		r.Form.Get("uid"),
		r.Form.Get("name"),
		r.Form.Get("email"),
		rating,
		r.Form.Get("comment"),
		notApprovedReview,
		time.Now()); err != nil {
		log.Println("broken sql;", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Response{
		OK:  true,
		Msg: "Review submitted, waiting for approval.",
	})
}

// Gravatar returns the URL to the image to identify an user.
func (app Application) Gravatar(email string) string {
	hasher := md5.New()
	hasher.Write([]byte(email))
	result := hex.EncodeToString(hasher.Sum(nil))
	return fmt.Sprintf("https://www.gravatar.com/avatar/%s", result)
}
