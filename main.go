package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/valyala/fasthttp"
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
	app.Get("/sse", events)
	app.Static("/", "./static")

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

func events(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		for {
			reading, err := getLatestReading()

			if err != nil {
				fmt.Printf("Error while getting reading data: %v", err)
				return
			}

			data, err := json.Marshal(reading)

			if err != nil {
				fmt.Printf("Error while marshaling json: %v", err)
				break
			}

			fmt.Printf("data: %s\n", string(data))
			fmt.Fprintf(w, "data: %s\n\n", string(data))

			err = w.Flush()
			if err != nil {
				// Refreshing page in web browser will establish a new
				// SSE connection, but only (the last) one is alive, so
				// dead connections must be closed here.
				fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
				break
			}
			time.Sleep(2 * time.Second)
		}
	}))

	return nil
}

func getLatestReading() (*Reading, error) {
	var reading Reading
	var id int
	row, err := db.Query("SELECT * FROM readings ORDER BY id DESC LIMIT 1")

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	row.Next()
	err = row.Scan(&id, &reading.Temperature, &reading.Humidity, &reading.DustConcentration, &reading.Pressure, &reading.AirPurity)

	if err != nil {
		log.Fatal(err)
	}

	return &reading, nil
}
