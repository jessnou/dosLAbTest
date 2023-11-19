package models

type MaxWord struct {
	PostID int    `db:"post_id"`
	Word   string `db:"word"`
	Count  int    `db:"count"`
}
