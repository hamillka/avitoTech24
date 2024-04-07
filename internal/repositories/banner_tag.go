package repositories

import "github.com/jmoiron/sqlx"

type BannerTagRepository struct {
	db *sqlx.DB
}

func NewBannerTagRepository(db *sqlx.DB) *BannerTagRepository {
	return &BannerTagRepository{db: db}
}
