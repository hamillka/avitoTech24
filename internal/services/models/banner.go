package models

import (
	"time"
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
