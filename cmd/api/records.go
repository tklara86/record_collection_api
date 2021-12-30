package main

import (
	"fmt"
	"net/http"
)

//  createRecordHandler creates new record
func (app *application) createRecordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a record")
}

// showRecordHandler shows specific record
func (app *application) showRecordHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of record %d", id)
}
