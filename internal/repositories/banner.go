package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/hamillka/avitoTech24/internal/handlers/dto"
	repoModels "github.com/hamillka/avitoTech24/internal/repositories/models"
	"github.com/jmoiron/sqlx"
)

type BannerRepository struct {
	db *sqlx.DB
}

const (
	selectBannersByFeatureForAdmin      = "SELECT * FROM banners WHERE feature_id = $1 LIMIT $2 OFFSET $3"
	selectBannersByFeatureForUser       = "SELECT * FROM banners WHERE feature_id = $1 AND is_active = true LIMIT $2 OFFSET $3"
	selectBannerByID                    = "SELECT * FROM banners WHERE banner_id = $1"
	selectBannerByFeatureAndTagForAdmin = "SELECT b.banner_id, b.feature_id, b.content, b.is_active, b.created_at, b.updated_at FROM banners b JOIN bt ON b.banner_id = bt.banner_id WHERE b.feature_id = $1 AND bt.tag_id = $2"
	selectBannerByFeatureAndTagForUser  = "SELECT b.banner_id, b.feature_id, b.content, b.is_active, b.created_at, b.updated_at FROM banners b JOIN bt ON b.banner_id = bt.banner_id WHERE b.is_active = true AND b.feature_id = $1 AND bt.tag_id = $2"
	selectTagsByBanner                  = "SELECT tag_id FROM bt WHERE banner_id = $1"
	selectBannersByTagForAdmin          = "SELECT banner_id FROM bt WHERE tag_id = $1 LIMIT $2 OFFSET $3"
	selectBannersByTagForUser           = "SELECT bt.banner_id FROM bt JOIN banners b ON bt.banner_id = b.banner_id WHERE tag_id = $1 AND b.is_active = true LIMIT $2 OFFSET $3"
	createBanner                        = "INSERT INTO banners (feature_id, content, is_active) VALUES ($1, $2, $3) RETURNING banner_id"
	updateBanner                        = "UPDATE banners SET feature_id = $1, content = $2, is_active = $3, updated_at = $4 WHERE banner_id = $5"
	deleteBannerByID                    = "DELETE FROM banners WHERE banner_id = $1"
)

func NewBannerRepository(db *sqlx.DB) *BannerRepository {
	return &BannerRepository{db: db}
}

func (br *BannerRepository) GetBannersByFeature(featureID, limit, offset, role int64) ([]*repoModels.Banner, map[int64][]int64, error) {
	var banners []*repoModels.Banner
	var (
		rows *sql.Rows
		err  error
	)

	if role == dto.ADMIN {
		rows, err = br.db.Query(selectBannersByFeatureForAdmin, featureID, limit, offset)
		if err != nil {
			return nil, nil, ErrDatabaseReadingError
		}
	} else if role == dto.USER {
		rows, err = br.db.Query(selectBannersByFeatureForUser, featureID, limit, offset)
		if err != nil {
			return nil, nil, ErrDatabaseReadingError
		}
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
		if err != nil {
			return nil, nil, ErrDatabaseReadingError
		}

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

func (br *BannerRepository) GetBannersByTag(tagID, limit, offset, role int64) ([]*repoModels.Banner, map[int64][]int64, error) {
	banners := make([]*repoModels.Banner, 0)
	var (
		rows *sql.Rows
		err  error
	)

	if role == dto.ADMIN {
		rows, err = br.db.Query(selectBannersByTagForAdmin, tagID, limit, offset)
		if err != nil {
			return nil, nil, ErrDatabaseReadingError
		}
	} else if role == dto.USER {
		rows, err = br.db.Query(selectBannersByTagForUser, tagID, limit, offset)
		if err != nil {
			return nil, nil, ErrDatabaseReadingError
		}
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
		err = br.db.QueryRow(selectBannerByID, bannerID).Scan(
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

func (br *BannerRepository) GetBannerByFeatureAndTag(featureID, tagID, role int64) (*repoModels.Banner, error) {
	banner := new(repoModels.Banner)
	var err error

	if role == dto.ADMIN {
		err = br.db.QueryRow(selectBannerByFeatureAndTagForAdmin, featureID, tagID).Scan(
			&banner.BannerID,
			&banner.FeatureID,
			&banner.Content,
			&banner.IsActive,
			&banner.CreatedAt,
			&banner.UpdatedAt,
		)
	} else if role == dto.USER {
		err = br.db.QueryRow(selectBannerByFeatureAndTagForUser, featureID, tagID).Scan(
			&banner.BannerID,
			&banner.FeatureID,
			&banner.Content,
			&banner.IsActive,
			&banner.CreatedAt,
			&banner.UpdatedAt,
		)
	}
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return banner, nil
}

func (br *BannerRepository) CreateBanner(featureID int64, content string, isActive bool) (int64, error) {
	var id int64

	row := br.db.QueryRow(createBanner, featureID, content, isActive)
	if err := row.Scan(&id); err != nil {
		return 0, ErrDatabaseWritingError
	}

	return id, nil
}

func (br *BannerRepository) UpdateBanner(bannerID, featureID int64, content string, isActive *bool) (int64, error) {
	banner := new(repoModels.Banner)

	err := br.db.QueryRow(selectBannerByID, bannerID).Scan(
		&banner.BannerID,
		&banner.FeatureID,
		&banner.Content,
		&banner.IsActive,
		&banner.CreatedAt,
		&banner.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrRecordNotFound
		}
		return 0, err
	}

	if featureID == 0 {
		featureID = banner.FeatureID
	}
	if content == "" {
		content = banner.Content
	}
	if isActive == nil {
		isActive = pointer.ToBool(banner.IsActive)
	}

	_, err = br.db.Exec(updateBanner,
		featureID,
		content,
		isActive,
		time.Now().Format(time.RFC3339),
		banner.BannerID,
	)
	if err != nil {
		return 0, ErrDatabaseUpdatingError
	}

	return banner.BannerID, nil
}

func (br *BannerRepository) DeleteBanner(bannerID int64) error {
	_, err := br.db.Exec(deleteBannerByID, bannerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrRecordNotFound
		}
		return err
	}

	return nil
}
