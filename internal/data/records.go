package data

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/tklara86/record_collection_api/internal/validator"
	"strconv"
	"time"
)

//WITH the_record AS (
//INSERT INTO records (title, label, year, cover) VALUES ('sdsd', 'sdsd', 1987, 'sdsdsd') RETURNING record_id
//)
//INSERT INTO record_genres (record_id, genre_id) VALUES
//((SELECT record_id from the_record), 1),
//((SELECT record_id from the_record), 2),
//((SELECT record_id from the_record), 3)

// Record model
type Record struct {
	RecordID     int64           `json:"record_id"`
	Title        string          `json:"title"`
	Label        string          `json:"label"`
	Year         int32           `json:"year"`
	Cover        string          `json:"cover"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	Genres       *[]Genres       `json:"genres,omitempty"`
	Artists      *[]Artists      `json:"artists,omitempty"`
	RecordGenre  *[]RecordGenre  `json:"record_genre,omitempty"`
	RecordArtist *[]RecordArtist `json:"record_artist,omitempty"`
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
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// RecordArtist model
type RecordArtist struct {
	ID        int64     `json:"id"`
	ArtistID  int64     `json:"artist_id"`
	RecordID  int64     `json:"record_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ValidateRecord(v *validator.Validator, record *Record) {
	v.Check(record.Title != "", "title", "must be provided")
	v.Check(len(record.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(record.Label != "", "label", "must be provided")
	v.Check(record.Year != 0, "year", "must be provided")
	v.Check(record.Cover != "", "cover", "must be provided")
}

// RecordModel a struct type which wraps a sql.DB connection pool.
type RecordModel struct {
	DB *sql.DB
}

// CreateRecord creates a new record in the record table
func (m RecordModel) CreateRecord(record *Record) error {
	q := `INSERT INTO records (title, label, year, cover) VALUES ($1, $2, $3, $4) RETURNING record_id`

	args := []interface{}{record.Title, record.Label, record.Year, record.Cover}

	return m.DB.QueryRow(q, args...).Scan(&record.RecordID)
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

// CreateGenreRecords
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

	fmt.Println(q)

	return m.DB.QueryRow(q, args...).Scan(recordGenre)
}
