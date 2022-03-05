package data

import "time"

// Artists model
type Artists struct {
	ArtistID  int64     `json:"artist_id"`
	Name      string    `json:"name"`
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
