package models

import "time"

type Banner struct {
	BannerID  int64     `db:"banner_id"`
	FeatureID int64     `db:"feature_id"`
	Content   string    `db:"content"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
