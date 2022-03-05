package data

import "time"

// Genres model
type Genres struct {
	GenreID   int64     `json:"genre_id"`
	GenreName string    `json:"genre_name"`
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
