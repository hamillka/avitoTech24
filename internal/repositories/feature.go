package repositories

import (
	"database/sql"
	"errors"

	repoModels "github.com/hamillka/avitoTech24/internal/repositories/models"
	"github.com/jmoiron/sqlx"
)

type FeatureRepository struct {
	db *sqlx.DB
}

const (
	selectFeatureByID = "SELECT * FROM features WHERE feature_id = $1"
)

func NewFeatureRepository(db *sqlx.DB) *FeatureRepository {
	return &FeatureRepository{db: db}
}

func (fr *FeatureRepository) GetFeatureByID(featureID int64) (*repoModels.Feature, error) {
	feature := new(repoModels.Feature)

	err := fr.db.QueryRow(selectFeatureByID, featureID).Scan(
		&feature.FeatureID,
		&feature.Name,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}

		return nil, err
	}

	return feature, nil
}
