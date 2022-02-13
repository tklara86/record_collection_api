package main

import (
	"errors"
	"fmt"
	"github.com/tklara86/record_collection_api/internal/data"
	"github.com/tklara86/record_collection_api/internal/validator"
	"net/http"
)

//INSERT INTO genres (genre_name) VALUES ('Classical'), ('Jazz'), ('Rock');
// createRecordHandler creates new record
func (app *application) createRecordHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		RecordID    int64              `json:"record_id"`
		Title       string             `json:"title"`
		Label       string             `json:"label"`
		Year        int32              `json:"year"`
		Cover       string             `json:"cover"`
		GenreID     int64              `json:"genre_id"`
		RecordGenre []data.RecordGenre `json:"record_genre"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	record := &data.Record{
		Title: input.Title,
		Label: input.Label,
		Year:  input.Year,
		Cover: input.Cover,
		RecordGenre: &[]data.RecordGenre{},
	}

	v := validator.New()

	if data.ValidateRecord(v, record, *record.RecordGenre); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Records.CreateRecord(record, *record.RecordGenre)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/records/%d", record.RecordID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"record": record}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

// showRecordHandler shows specific record
func (app *application) showRecordHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	record, err := app.models.Records.GetRecord(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrorRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"record": record}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
