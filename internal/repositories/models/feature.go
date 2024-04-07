package models

type Feature struct {
	FeatureID int64  `db:"feature_id"`
	Name      string `db:"name"`
}
