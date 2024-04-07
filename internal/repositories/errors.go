package repositories

import (
	goErrors "errors"
)

var (
	ErrRecordNotFound       = goErrors.New("Record was not found")
	ErrDatabaseWritingError = goErrors.New("Error while writing to DB")
	ErrDatabaseReadingError = goErrors.New("Error while reading from DB")
	ErrRecordAlreadyExists  = goErrors.New("Record with this data already exists")
)
