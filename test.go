package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"strings"

	// Importing gorilla/sessions package for session management
	"github.com/gorilla/sessions"

	// Importing go-sqlite3 driver for using SQLite as the database
	_ "github.com/mattn/go-sqlite3"

	// Importing bcrypt package for password hashing
	"golang.org/x/crypto/bcrypt"
)

// Declaring the global session store variable
var store = sessions.NewCookieStore([]byte("secret-key"))

// Defining the NewsItem struct
type NewsItem struct {
	Id          int
	Title       string
	Text        string
	Date        string
}

// Defining the FeedbackItem struct
type FeedbackItem struct {
	Id          int
	Name        string
	Email       string
	Text        string
}

// Defining the SupItem struct
type SupItem struct {
	Id          int
	Image       string
	Name        string
	Short       string
	Long        string
}

// Defining the RentItem struct
type RentItem struct {
	Id          int
	Name        string
	Square      int
	Price       int
	Image       string
	Phone       string
	Email       string
	About       string
	Ind         int
	InBasket    string
}

// Defining the SaleItem struct
type SaleItem struct {
	Id          int
	Title       string
	Text        string
}

// registr function handles user registration
func registr(w http.ResponseWriter, r *http.Request) {
	// Opening SQLite database connection
	db, err := sql.Open("sqlite3", "my_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Getting session for the user
	session, err := store.Get(r, "user-session")
	_, ok := session.Values["user"]
	if ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else if r.Method == http.MethodPost {
		// Parsing and processing the registration form
	} else {
		// Displaying the registration form
	}
}

// index function handles the homepage
func index(w http.ResponseWriter, r *http.Request) {
	// Parsing and rendering the index template
}

// login function handles user login
func login(w http.ResponseWriter, r *http.Request) {
	// Opening SQLite database connection
	db, err := sql.Open("sqlite3", "my_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Getting session for the user
	session, err := store.Get(r, "user-session")
	_, ok := session.Values["user"]
	if ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else if r.Method == http.MethodPost {
		// Parsing and processing the login form
	} else {
		// Displaying the login form
	}
}

// logout function handles user logout
func logout(w http.ResponseWriter, r *http.Request) {
	// Getting session for the user
	session, err := store.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Setting session expiration to -1 to destroy the session
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirecting the user to the homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ... other functions follow the same pattern with comments

func main() {
	// Starting the HTTP server
	handle_request()
}
