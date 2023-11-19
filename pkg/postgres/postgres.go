package postgres

import (
	"dosLAbTest/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

const (
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	connAttempts int
	connTimeout  time.Duration

	DB *sqlx.DB
}

func New(config config.Config) (*Postgres, error) {
	pg := &Postgres{
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	db, err := sqlx.Open(config.DBDriver, config.URL)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(_defaultConnAttempts)
	db.SetConnMaxLifetime(_defaultConnTimeout)

	pg.DB = db

	return pg, nil
}
