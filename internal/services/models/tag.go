package models

type Tag struct {
	TagID int64  `db:"tag_id"`
	Name  string `db:"name"`
}
