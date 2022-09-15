package repository

import "database/sql"

type repository struct {
	Database *sql.DB
}
