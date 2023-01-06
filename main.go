package main

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

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

	app := fiber.New()

	app.Post("/api/readings/add", addReading)

	app.Listen(":4000")

	log.Fatal("Failed to start server")
}

// TODO: Move to its own model package
type AirPurity string

const (
	Dangerous AirPurity = "dangerous"
	High      AirPurity = "high"
	Low       AirPurity = "low"
	FreshAir  AirPurity = "fresh_air"
)

type Reading struct {
	Temperature       float32   `json:"temperature"`
	Humidity          float32   `json:"humidity"`
	DustConcentration float32   `json:"dust_concentration"`
	Pressure          int32     `json:"pressure"`
	AirPurity         AirPurity `json:"air_purity"`
}

// TODO: Move to its own handler package
func addReading(c *fiber.Ctx) error {
	var reading Reading
	if err := c.BodyParser(&reading); err != nil {
		return err
	}

	stmt := `INSERT INTO readings (temperature, humidity, dust_concentration, pressure, air_purity) 
	VALUES($1, $2, $3, $4, $5) RETURNING id`

	var id int
	err := db.QueryRow(stmt, reading.Temperature, reading.Humidity, reading.DustConcentration, reading.Pressure, reading.AirPurity).Scan(&id)
	if err != nil {
		return err
	}

	c.SendString("Record inserted: " + strconv.Itoa(id))

	return nil
}
