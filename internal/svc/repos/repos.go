package repos

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repos struct {
	db               *sqlx.DB
	AttributeDefRepo *AttributeDefRepo
}

func NewRepos(
	driver string,
	dataSourceName string,
	maxOpenConns int,
	maxIdleConns int,
	maxIdleTime string,
) (*Repos, error) {

	connMaxIdleTime, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening postgres db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to postgres db: %w", err)
	}
	db.DB.SetMaxOpenConns(maxOpenConns)
	db.DB.SetMaxIdleConns(maxIdleConns)
	db.DB.SetConnMaxIdleTime(connMaxIdleTime)

	return &Repos{
		db:               db,
		AttributeDefRepo: &AttributeDefRepo{db: db},
	}, nil
}

// Stop closes the database connections.
func (r *Repos) Stop() {
	if err := r.db.Close(); err != nil {
		slog.Warn("Failed to close connections.", "error", err)
	} else {
		slog.Info("Database connections closed.")
	}
}
