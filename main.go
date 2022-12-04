package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/naconnors/ambi-go/models"
)

type application struct {
	readings *models.ReadingModel
}

// TODO: Use global database for now, this should be changed later
var db *sql.DB

func main() {

	var err error
	//TODO: This should be moved to its own package
	db, err = sql.Open("pgx", "postgres://postgres:postgres@localhost/ambi_go_dev")
	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Unable to ping database")
	}

	defer db.Close()

	app := &application{
		readings: &models.ReadingModel{DB: db},
	}

	server := &http.Server{
		Addr:    "localhost:4000",
		Handler: app.routes(),
	}

	err = server.ListenAndServe()

	log.Fatalf("Server failed to start: %v", err)
}
