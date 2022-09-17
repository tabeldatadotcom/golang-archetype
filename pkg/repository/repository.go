package repository

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	Connect *sqlx.DB
}
