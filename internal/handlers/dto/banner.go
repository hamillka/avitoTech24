package dto

import (
	"time"

	"github.com/hamillka/avitoTech24/internal/services/models"
)

type CreateOrUpdateBannerRequestDto struct {
	TagIDs    []int64 `json:"tag_ids"`
	FeatureID int64   `json:"feature_id"`
	Content   string  `json:"content"`
	IsActive  *bool   `json:"is_active"`
}

type CreateOrUpdateBannerResponseDto struct {
	ID int64 `json:"id"`
}

type GetBannersResponseDto struct {
	BannerID  int64     `json:"banner_id"`
	FeatureID int64     `json:"feature_id"`
	TagIDs    []int64   `json:"tag_ids"`
	Content   string    `json:"content"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserBannerResponseDto struct {
	Content string `json:"content"`
}

func ConvertToDto(banner models.BannerWithTagIDs) *GetBannersResponseDto {
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
