package repos

import (
	"log/slog"

	"github.com/dxps/tmc-pwa/internal/shared/model/meta"
	"github.com/jmoiron/sqlx"
)

type AttributeDefRepo struct {
	db *sqlx.DB
}

func NewAttributeDefRepo() *AttributeDefRepo {
	return &AttributeDefRepo{}
}

func (r *AttributeDefRepo) GetAll() ([]meta.AttributeDef, error) {
	slog.Debug("Getting all attribute definitions ...")
	_ = r.db.Ping()
	return []meta.AttributeDef{}, nil
}
