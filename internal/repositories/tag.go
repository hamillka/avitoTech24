package repositories

import "github.com/jmoiron/sqlx"

type TagRepository struct {
	db *sqlx.DB
}

func NewTagRepository(db *sqlx.DB) *TagRepository {
	return &TagRepository{db: db}
}
