package repositories

import (
	repoModels "github.com/hamillka/avitoTech24/internal/repositories/models"
	"github.com/jmoiron/sqlx"
)

type BannerRepository struct {
	db *sqlx.DB
}

const (
	selectBannersByFeature = "SELECT * FROM banners WHERE feature_id = $1"
	selectTagsByBanner     = "SELECT tag_id FROM bt WHERE banner_id = $1"
)

func NewBannerRepository(db *sqlx.DB) *BannerRepository {
	return &BannerRepository{db: db}
}

func (br *BannerRepository) GetBannersByFeature(featureID, limit, offset int64) ([]*repoModels.Banner, map[int64][]int64, error) {
	var banners []*repoModels.Banner

	rows, err := br.db.Query(selectBannersByFeature, featureID)
	if err != nil {
		return nil, nil, ErrDatabaseReadingError
	}

	if err := rows.Err(); err != nil {
		return nil, nil, ErrDatabaseReadingError
	}

	defer rows.Close()

	for rows.Next() {
		banner := new(repoModels.Banner)
		if err := rows.Scan(&banner.BannerID, &banner.FeatureID, &banner.Content, &banner.IsActive, &banner.CreatedAt, &banner.UpdatedAt); err != nil {
			return nil, nil, ErrDatabaseReadingError
		}
		banners = append(banners, banner)
	}

	bannerTagMap := make(map[int64][]int64)
	for _, banner := range banners {
		var tagIDs []int64

		rows, err = br.db.Query(selectTagsByBanner, banner.BannerID)

		for rows.Next() {
			var tagID int64
			if err := rows.Scan(&tagID); err != nil {
				return nil, nil, ErrDatabaseReadingError
			}
			tagIDs = append(tagIDs, tagID)
		}
		bannerTagMap[banner.BannerID] = tagIDs
	}

	return banners, bannerTagMap, nil
}

func (br *BannerRepository) GetBannersByTag(tagId, limit, offset int64) ([]*repoModels.Banner, map[int64][]int64, error) {
	return nil, nil, nil
}

func (br *BannerRepository) GetBannerByFeatureAndTag(featureId, tagId, limit, offset int64) (*repoModels.Banner, error) {
	return nil, nil
}
