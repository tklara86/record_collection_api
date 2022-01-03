package main

import (
	json2 "encoding/json"
	"fmt"
	"github.com/tklara86/record_collection_api/internal/data"
	"net/http"
	"time"
)

//  createRecordHandler creates new record
func (app *application) createRecordHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		RecordID int64         `json:"record_id"`
		Title    string        `json:"title"`
		Label    string        `json:"label"`
		Year     int32         `json:"year"`
		Cover    string        `json:"cover"`
		Genres   []interface{} `json:"genres,omitempty"`
	}

	err := json2.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
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
