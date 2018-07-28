package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/stephenhillier/soildesc/backend/models"
)

// Server represents the server environment (db and router)
type Server struct {
	db     models.Datastore
	router chi.Router
}

func main() {

	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	dbname := os.Getenv("DBNAME")
	dbhost := os.Getenv("DBHOST")

	// create db connection and router and use them to create a new "Server" instance
	db, err := models.NewDB(fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbuser, dbpass, dbhost, dbname))
	if err != nil {
		log.Panic(err)
	}
	r := chi.NewRouter()
	s := &Server{db, r}

	// register middleware
	s.router.Use(middleware.Logger)

	// register routes from routes.go
	s.routes()

	log.Printf("Starting HTTP server on port 8000.\n")
	log.Printf("Press CTRL+C to stop.")
	log.Fatal(http.ListenAndServe(":8000", s.router))
}
