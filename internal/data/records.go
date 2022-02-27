package data

import (
	"database/sql"
	"errors"
	"github.com/tklara86/record_collection_api/internal/validator"
	"strconv"
	"time"
)

// Record model
type Record struct {
	RecordID      int64          `json:"record_id"`
	Title         string         `json:"title"`
	Label         string         `json:"label"`
	Year          int32          `json:"year"`
	Cover         string         `json:"cover"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	RecordGenres  []RecordGenre  `json:"record_genres,omitempty"`
	RecordArtists []RecordArtist `json:"record_artist,omitempty"`
}

// Genres model
type Genres struct {
	GenreID   int64     `json:"genre_id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Artists model
type Artists struct {
	ArtistID  int64     `json:"artist_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// RecordGenre model
type RecordGenre struct {
	ID        int64     `json:"id"`
	RecordID  int64     `json:"record_id"`
	GenreID   int64     `json:"genre_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// RecordArtist model
type RecordArtist struct {
	ID        int64     `json:"id"`
	ArtistID  int64     `json:"artist_id"`
	RecordID  int64     `json:"record_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ValidateRecord validates record fields
func ValidateRecord(v *validator.Validator, record *Record, recordGenre []RecordGenre,
	recordArtist []RecordArtist) {
	v.Check(record.Title != "", "title", "must be provided")
	v.Check(len(record.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(record.Label != "", "label", "must be provided")
	v.Check(record.Year != 0, "year", "must be provided")
	v.Check(record.Cover != "", "cover", "must be provided")

	if len(recordGenre) < 1 {
		v.AddError("genre", "Please select genre(s)")
	}
	if len(recordArtist) < 1 {
		v.AddError("artist", "Please select artist(s)")
	}

}

// RecordModel a struct type which wraps a sql.DB connection pool.
type RecordModel struct {
	DB *sql.DB
}

// CreateRecord creates a new record in the record table
func (m RecordModel) CreateRecord(record *Record, recordGenre []RecordGenre,
	artist []RecordArtist) error {
	q := `WITH the_record AS (
    	INSERT INTO records (title, label, year, cover) VALUES ($1, $2, $3, $4) RETURNING record_id
	),
	genre AS (INSERT INTO record_genres (record_id, genre_id) VALUES `

	args := []interface{}{record.Title, record.Label, record.Year, record.Cover}

	var nc int

	for i, v := range recordGenre {
		args = append(args, v.GenreID)
		numFields := 1
		nc = (i * numFields) + 4

		for j := 0; j < numFields; j++ {
			q += `((SELECT record_id from the_record),` + `$` + strconv.Itoa(nc+j+1) + `),`
		}
		q = q[:len(q)-1] + `,`
	}
	q = q[:len(q)-1]

	q += ` RETURNING record_id)`

	q += ` INSERT INTO record_artists (record_id, artist_id) VALUES `

	for index, j := range artist {
		args = append(args, j.ArtistID)
		numFields := 1
		nk := (index * numFields) + nc + 1

		for k := 0; k < numFields; k++ {
			q += `((SELECT record_id from the_record),` + `$` + strconv.Itoa(nk+k+1) + `),`
		}
		q = q[:len(q)-1] + `,`
	}
	q = q[:len(q)-1]

	q += ` RETURNING record_id`

	return m.DB.QueryRow(q, args...).Scan(&record.RecordID)
}

// CreateGenreRecords creates genre records in the record_genres table
func (m RecordModel) CreateGenreRecords(recordGenre []RecordGenre) error {
	q := `INSERT INTO record_genres(record_id,genre_id) VALUES `

	var args []interface{}

	for i, v := range recordGenre {
		args = append(args, v.RecordID, v.GenreID)
		numFields := 2
		n := i * numFields

		q += `(`
		for j := 0; j < numFields; j++ {
			q += `$` + strconv.Itoa(n+j+1) + `,`
		}
		q = q[:len(q)-1] + `),`

	}
	q = q[:len(q)-1]

	return m.DB.QueryRow(q, args...).Scan(recordGenre)
}

// GetRecord fetches specific record from the record table
func (m RecordModel) GetRecord(id int64) (*Record, error) {
	q := `SELECT * FROM 
          records
          WHERE record_id = $1`

	var record Record

	err := m.DB.QueryRow(q, id).Scan(&record.RecordID, &record.Title, &record.Label,
		&record.Year, &record.Cover, &record.CreatedAt, &record.UpdatedAt)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorRecordNotFound
		default:
			return nil, err
		}
	}

	return &record, nil
}

// UpdateRecord updates specific record in the record table
func (m RecordModel) UpdateRecord(record *Record) error {
	return nil
}

// DeleteRecord deletes specific record
func (m RecordModel) DeleteRecord(id int64) error {
	return nil
}
