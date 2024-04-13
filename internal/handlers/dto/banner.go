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
	Content   map[string]interface{} `json:"content"`    // Контент баннера
	IsActive  *bool                  `json:"is_active"`  // Статус активности баннера
	TagIDs    []int64                `json:"tag_ids"`    // Теги баннера
	FeatureID int64                  `json:"feature_id"` // Фича баннера
}

// CreateOrUpdateBannerResponseDto model info
// @Description Информация о баннере при создании или изменении
type CreateOrUpdateBannerResponseDto struct {
	ID int64 `json:"id"` // Идентификатор баннера
}

// GetBannersResponseDto model info
// @Description Информация о баннере при получении баннеров
type GetBannersResponseDto struct {
	BannerID  int64     `json:"banner_id"`  // Идентификатор баннера
	FeatureID int64     `json:"feature_id"` // Фича баннера
	TagIDs    []int64   `json:"tag_ids"`    // Теги баннера
	Content   string    `json:"content"`    // Контент баннера
	IsActive  bool      `json:"is_active"`  // Статус активности баннера
	CreatedAt time.Time `json:"created_at"` // Время создания баннера
	UpdatedAt time.Time `json:"updated_at"` // Время обновления баннера
}

// GetUserBannerResponseDto model info
// @Description Информация о контенте при получении баннера
type GetUserBannerResponseDto struct {
	Content string `json:"content"` // Контент баннера
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
