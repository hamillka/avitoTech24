package dto

import (
	"time"

	"github.com/hamillka/avitoTech24/internal/services/models"
)

const (
	ADMIN = iota
	USER
)

// CreateOrUpdateBannerRequestDto model info
// @Description Информация о баннере при создании или изменении
type CreateOrUpdateBannerRequestDto struct {
	Content   map[string]interface{} `json:"content"`
	IsActive  *bool                  `json:"is_active"`
	TagIDs    []int64                `json:"tag_ids"`
	FeatureID int64                  `json:"feature_id"`
}

// CreateOrUpdateBannerResponseDto model info
// @Description Информация о баннере при создании или изменении
type CreateOrUpdateBannerResponseDto struct {
	ID int64 `json:"id"`
}

// GetBannersResponseDto model info
// @Description Информация о баннере при получении баннеров
type GetBannersResponseDto struct {
	BannerID  int64     `json:"banner_id"`
	FeatureID int64     `json:"feature_id"`
	TagIDs    []int64   `json:"tag_ids"`
	Content   string    `json:"content"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetUserBannerResponseDto model info
// @Description Информация о контенте при получении баннера
type GetUserBannerResponseDto struct {
	Content string `json:"content"`
}

func ConvertToDto(banner *models.BannerWithTagIDs) *GetBannersResponseDto {
	return &GetBannersResponseDto{
		BannerID:  banner.BannerID,
		FeatureID: banner.FeatureID,
		TagIDs:    banner.TagIDs,
		Content:   banner.Content,
		IsActive:  banner.IsActive,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
	}
}
