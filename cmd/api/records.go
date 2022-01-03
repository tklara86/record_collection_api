package main

import (
	"fmt"
	"github.com/tklara86/record_collection_api/internal/data"
	"github.com/tklara86/record_collection_api/internal/validator"
	"net/http"
	"time"
)

//  createRecordHandler creates new record
func (app *application) createRecordHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		RecordID    int64              `json:"record_id"`
		Title       string             `json:"title"`
		Label       string             `json:"label"`
		Year        int32              `json:"year"`
		Cover       string             `json:"cover"`
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
		RecordGenre: &[]data.RecordGenre{
			{
				ID:       2,
				RecordID: 1,
				GenreID:  1,
			},
			{
				ID:       3,
				RecordID: 1,
				GenreID:  2,
			},
		},
	}

	v := validator.New()

	if data.ValidateRecord(v, record); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)

}

// showRecordHandler shows specific record
func (app *application) showRecordHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	record := &data.Record{
		RecordID:  id,
		Title:     "Actions",
		Label:     "DG",
		Year:      1978,
		Cover:     "actions.jpg",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Genres: &[]data.Genres{
			{
				GenreID:   1,
				GenreName: "Jazz",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				GenreID:   2,
				GenreName: "Classical",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		Artists: &[]data.Artists{
			{
				ArtistID:  3,
				FirstName: "Krzysztof",
				LastName:  "Penderecki",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ArtistID:  2,
				FirstName: "Don",
				LastName:  "Chery",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"record": record}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
