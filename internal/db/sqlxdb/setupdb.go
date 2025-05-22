package sqlxdb

import (
	"context"

	"database/sql"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pricetula/gaze-news-api/internal/utils"
)

func SetupDB(ctx context.Context, cfg *utils.Config) (db *sqlx.DB, err error) {
	// Open a new connection to the database
	sqldb, err := sql.Open("postgres", cfg.DB.Connection)
	if err != nil {
		return
	}

	// Run the migrations
	err = dbmigrateup(sqldb)
	if err != nil {
		return
	}

	// Create a new sqlx db instance
	db = sqlx.NewDb(sqldb, "postgres")
	return
}
