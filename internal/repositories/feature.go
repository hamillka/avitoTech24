package repositories

import "github.com/jmoiron/sqlx"

type FeatureRepository struct {
	db *sqlx.DB
}

func NewFeatureRepository(db *sqlx.DB) *FeatureRepository {
	return &FeatureRepository{db: db}
}
