package query

import (
	"dosLAbTest/pkg/postgres"
	"dosLAbTest/pkg/postgres/models"
	"dosLAbTest/pkg/postgres/transaction"
	"github.com/jmoiron/sqlx"
)

const (
	upsertPost = `INSERT INTO posts (post_id, word, count) VALUES (:post_id,:word,:count) 
					ON CONFLICT (post_id,word) DO UPDATE SET 
					count = Excluded.count`
	getPost = `SELECT post_id,word,count FROM posts WHERE post_id = $1 ORDER BY count DESC `
)

func Insert(postgres postgres.Postgres, max []models.MaxWord) error {
	return transaction.WithTransaction(postgres.DB, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExec(upsertPost, max)
		if err != nil {
			return err
		}

		return nil
	})
}

func GetPost(postgres postgres.Postgres, id int) ([]models.MaxWord, error) {
	var word []models.MaxWord
	err := postgres.DB.Select(&word, getPost, id)
	if err != nil {
		return nil, err
	}
	return word, nil
}
