package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/julienschmidt/httprouter"
)

// Use global database for now, this should be changed later
var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("pgx", "postgres://postgres:postgres@localhost/ambi_go_dev")
	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Unable to ping database")
	}

	defer db.Close()

	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/readings/add", addReading)

	log.Fatal(http.ListenAndServe(":4000", router))
}

type AirPurity string

const (
	Dangerous AirPurity = "dangerous"
	High      AirPurity = "high"
	Low       AirPurity = "low"
	FreshAir  AirPurity = "fresh_air"
)

type Reading struct {
	Temperature       float32
	Humidity          float32
	DustConcentration float32
	Pressure          int32
	AirPurity         AirPurity
}

func addReading(w http.ResponseWriter, r *http.Request) {
	var reading Reading
	err := json.NewDecoder(r.Body).Decode(&reading)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt := `INSERT INTO readings (temperature, humidity, dust_concentration, pressure, air_purity) 
	VALUES($1, $2, $3, $4, $5) RETURNING id`

	var id int
	err = db.QueryRow(stmt, reading.Temperature, reading.Humidity, reading.DustConcentration, reading.Pressure, reading.AirPurity).Scan(&id)
	if err != nil {
		log.Fatal("Error inserting reading into database")
	}

	fmt.Fprintf(w, "Record inserted: %v", id)
}
