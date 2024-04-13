package repositories

import (
	"errors"
)

var (
	ErrRecordNotFound        = errors.New("record was not found")
	ErrDatabaseWritingError  = errors.New("error while writing to DB")
	ErrDatabaseReadingError  = errors.New("error while reading from DB")
	ErrDatabaseUpdatingError = errors.New("record was not updated")
)
