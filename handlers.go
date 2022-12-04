package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/naconnors/ambi-go/models"
)

func (app *application) addReading(w http.ResponseWriter, r *http.Request) {
	var reading models.Reading
	err := json.NewDecoder(r.Body).Decode(&reading)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := app.readings.Insert(reading)
	if err != nil {
		log.Fatalf("inserting reading record: %v", err)
	}

	fmt.Fprintf(w, "Record inserted: %v", id)
}
