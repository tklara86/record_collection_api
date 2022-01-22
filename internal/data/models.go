package data

import (
	"database/sql"
	"errors"
)

var (
	ErrorRecordNotFound = errors.New("record not found")
)

type Models struct {
	Records RecordModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Records: RecordModel{
			DB: db,
		},
	}
}
