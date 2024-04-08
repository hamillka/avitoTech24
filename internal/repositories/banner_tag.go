package repositories

import "github.com/jmoiron/sqlx"

type BannerTagRepository struct {
	db *sqlx.DB
}

const (
	createBannerTag = "INSERT INTO bt (banner_id, tag_id) VALUES ($1, $2)"
)

func NewBannerTagRepository(db *sqlx.DB) *BannerTagRepository {
	return &BannerTagRepository{db: db}
}

func (btr *BannerTagRepository) CreateBannerTag(bannerID, tagID int64) error {
	_, err := btr.db.Exec(createBannerTag, bannerID, tagID)
	if err != nil {
		return ErrDatabaseWritingError
	}

	return nil
}
