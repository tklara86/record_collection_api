package main

import (
	"errors"
	"fmt"
	"github.com/tklara86/record_collection_api/internal/data"
	"github.com/tklara86/record_collection_api/internal/validator"
	"net/http"
)

// createRecordHandler creates new record
func (app *application) createRecordHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		RecordID      int64               `json:"record_id"`
		Title         string              `json:"title"`
		Release       string              `json:"release"`
		Cover         string              `json:"cover"`
		Status        string              `json:"status"`
		RecordGenres  []data.RecordGenre  `json:"record_genres"`
		RecordArtists []data.RecordArtist `json:"record_artists"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var recordGenres []data.RecordGenre
	var recordArtists []data.RecordArtist

	recordGenres = append(recordGenres, input.RecordGenres...)
	recordArtists = append(recordArtists, input.RecordArtists...)

	record := &data.Record{
		Title:         input.Title,
		Release:       input.Release,
		Cover:         input.Cover,
		Status:        input.Status,
		RecordGenres:  recordGenres,
		RecordArtists: recordArtists,
	}

	v := validator.New()

	if data.ValidateRecord(v, record, record.RecordGenres, record.RecordArtists); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Records.CreateRecord(record, record.RecordGenres, record.RecordArtists)
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
	artists, err := app.models.Records.GetRecordArtists(id)
	genres, err := app.models.Records.GetRecordGenres(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrorRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	for _, a := range artists {
		ra := &data.Artists{
			Name: a.Name,
		}

		record.Artists = append(record.Artists, *ra)
	}

	for _, g := range genres {
		rg := &data.Genres{
			GenreName: g.GenreName,
		}

		record.Genres = append(record.Genres, *rg)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"record": record}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) showRecordsHandler(w http.ResponseWriter, r *http.Request) {
	records, err := app.models.Records.GetAllRecords()

	if err != nil {
		switch {
		case errors.Is(err, data.ErrorRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"records": records}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
