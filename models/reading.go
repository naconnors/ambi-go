package models

import "database/sql"

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

type ReadingModel struct {
	DB *sql.DB
}

func (r *ReadingModel) Insert(reading Reading) (*Reading, error) {

	stmt := `INSERT INTO readings (temperature, humidity, dust_concentration, pressure, air_purity) 
	VALUES($1, $2, $3, $4, $5) RETURNING id`

	var id int
	err := r.DB.QueryRow(stmt, reading.Temperature, reading.Humidity, reading.DustConcentration, reading.Pressure, reading.AirPurity).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &reading, nil
}
