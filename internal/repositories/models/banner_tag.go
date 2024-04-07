package models

type BannerTag struct {
	BannerID int64 `db:"banner_id"`
	TagID    int64 `db:"tag_id"`
}
