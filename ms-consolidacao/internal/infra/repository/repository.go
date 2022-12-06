package repository

import (
	"database/sql"

	"github.com/yants95/imersao-full-cycle-cartola-fc/internal/infra/db"
)

type Repository struct {
	dbConn *sql.DB
	*db.Queries
}
