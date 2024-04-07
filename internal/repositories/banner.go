package repositories

import (
	repoModels "github.com/hamillka/avitoTech24/internal/repositories/models"
	"github.com/jmoiron/sqlx"
)

type BannerRepository struct {
	db *sqlx.DB
}

const (
	selectBannersByFeature      = "SELECT * FROM banners WHERE feature_id = $1 LIMIT $2 OFFSET $3"
	selectBannersByID           = "SELECT * FROM banners WHERE banner_id = $1 LIMIT $2 OFFSET $3"
	selectBannerByFeatureAndTag = "SELECT (b.banner_id, b.feature_id, b.content, b.is_active, b.created_at, b.updated_at) FROM banners b JOIN bt ON b.banner_id = bt.banner_id WHERE b.feature_id = $1 AND bt.tag_id = $2 LIMIT $3 OFFSET $4"
	selectTagsByBanner          = "SELECT tag_id FROM bt WHERE banner_id = $1"
	selectBannersByTag          = "SELECT banner_id FROM bt WHERE tag_id = $1"
)

func NewBannerRepository(db *sqlx.DB) *BannerRepository {
	return &BannerRepository{db: db}
}

func (br *BannerRepository) GetBannersByFeature(featureID, limit, offset int64) ([]*repoModels.Banner, map[int64][]int64, error) {
	var banners []*repoModels.Banner

	rows, err := br.db.Query(selectBannersByFeature, featureID, limit, offset)
	if err != nil {
		return nil, nil, ErrDatabaseReadingError
	}

	if err := rows.Err(); err != nil {
		return nil, nil, ErrDatabaseReadingError
	}

	defer rows.Close()

	for rows.Next() {
		banner := new(repoModels.Banner)
		if err := rows.Scan(
			&banner.BannerID,
			&banner.FeatureID,
			&banner.Content,
			&banner.IsActive,
			&banner.CreatedAt,
			&banner.UpdatedAt,
		); err != nil {
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

func (br *BannerRepository) GetBannersByTag(tagID, limit, offset int64) ([]*repoModels.Banner, map[int64][]int64, error) {
	var banners []*repoModels.Banner

	rows, err := br.db.Query(selectBannersByTag, tagID)
	if err != nil {
		return nil, nil, ErrDatabaseReadingError
	}

	if err := rows.Err(); err != nil {
		return nil, nil, ErrDatabaseReadingError
	}

	defer rows.Close()

	var bannerIDs []int64
	for rows.Next() {
		var bannerID int64
		if err := rows.Scan(&bannerID); err != nil {
			return nil, nil, ErrDatabaseReadingError
		}
		bannerIDs = append(bannerIDs, bannerID)
	}

	bannerTagMap := make(map[int64][]int64)

	for _, bannerID := range bannerIDs {
		bannerTagMap[bannerID] = []int64{tagID}

		banner := new(repoModels.Banner)
		err = br.db.QueryRow(selectBannersByID, bannerID).Scan(
			&banner.BannerID,
			&banner.FeatureID,
			&banner.Content,
			&banner.IsActive,
			&banner.CreatedAt,
			&banner.UpdatedAt,
		)
		if err != nil {
			return nil, nil, ErrDatabaseReadingError
		}
		banners = append(banners, banner)
	}

	return banners, bannerTagMap, nil
}

func (br *BannerRepository) GetBannerByFeatureAndTag(featureID, tagID, limit, offset int64) (*repoModels.Banner, error) {
	banner := new(repoModels.Banner)

	err := br.db.QueryRow(selectBannerByFeatureAndTag, featureID, tagID, limit, offset).Scan(
		&banner.BannerID,
		&banner.FeatureID,
		&banner.Content,
		&banner.IsActive,
		&banner.CreatedAt,
		&banner.UpdatedAt,
	)

	if err != nil {
		return nil, ErrRecordNotFound
	}

	return banner, nil
}
