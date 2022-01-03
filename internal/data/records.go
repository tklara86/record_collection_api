package data

import "time"

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
