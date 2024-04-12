package models

import (
	"time"

	"github.com/hamillka/avitoTech24/internal/repositories/models"
)

type BannerWithTagIDs struct {
	BannerID  int64
	FeatureID int64
	TagIDs    []int64
	Content   string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ConvertToBL(banner *models.Banner, tagIDs map[int64][]int64) *BannerWithTagIDs {
	return &BannerWithTagIDs{
		BannerID:  banner.BannerID,
		FeatureID: banner.FeatureID,
		TagIDs:    tagIDs[banner.BannerID],
		Content:   banner.Content,
		IsActive:  banner.IsActive,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
	}
}
