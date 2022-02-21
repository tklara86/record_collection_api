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
		RecordID      int64               `json:"record_id"`
		Title         string              `json:"title"`
		Label         string              `json:"label"`
		Year          int32               `json:"year"`
		Cover         string              `json:"cover"`
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

	// mock the ids from form
	genreSlices := []int64{1, 2, 3}

	var genres []data.RecordGenre
	var artists []data.RecordArtist

	for _, genreId := range genreSlices {

		d := []data.RecordGenre{
			{
				GenreID: genreId,
			},
		}

		a := []data.RecordArtist{
			{
				ArtistID: genreId,
			},
		}
		genres = append(genres, d...)
		artists = append(artists, a...)
	}

	record := &data.Record{
		Title:         input.Title,
		Label:         input.Label,
		Year:          input.Year,
		Cover:         input.Cover,
		RecordGenres:  genres,
		RecordArtists: artists,
	}

	v := validator.New()

	if data.ValidateRecord(v, record, record.RecordGenres); !v.Valid() {
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
