package repositories

import (
	"database/sql"
	"errors"

	repoModels "github.com/hamillka/avitoTech24/internal/repositories/models"
	"github.com/jmoiron/sqlx"
)

type TagRepository struct {
	db *sqlx.DB
}

const (
	selectTagByID = "SELECT * FROM tags WHERE tag_id = $1"
)

func NewTagRepository(db *sqlx.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (tr *TagRepository) GetTagByID(tagID int64) (*repoModels.Tag, error) {
	tag := new(repoModels.Tag)

	err := tr.db.QueryRow(selectTagByID, tagID).Scan(
		&tag.TagID,
		&tag.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}

		return nil, err
	}

	return tag, nil
}
