package models

type User struct {
	UserID  int64  `db:"user_id"`
	Token   string `db:"token"`
	IsAdmin bool   `db:"is_admin"`
}
