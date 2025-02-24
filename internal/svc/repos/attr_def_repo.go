package repos

import (
	"fmt"
	"log/slog"

	"github.com/dxps/tmc-pwa/internal/shared/model/meta"
	"github.com/jmoiron/sqlx"
)

const (
	getAllAttributeDefsSql = `SELECT * FROM attribute_defs`
	addAttributeDefSql     = `INSERT INTO attribute_defs (id, name, description, value_type, default_value, required) 
	                          VALUES (:id, :name, :description, :value_type, :default_value, :required)`
)

type AttributeDefRepo struct {
	db *sqlx.DB
}

func NewAttributeDefRepo() *AttributeDefRepo {
	return &AttributeDefRepo{}
}

func (r *AttributeDefRepo) Add(entry *meta.AttributeDef) error {
	_, err := r.db.NamedExec(addAttributeDefSql, entry)
	return err
}

func (r *AttributeDefRepo) GetAll() ([]*meta.AttributeDef, error) {
	entries := []*meta.AttributeDef{}
	if err := r.db.Select(&entries, getAllAttributeDefsSql); err != nil {
		return nil, err
	}
	slog.Debug(fmt.Sprintf("Got %d entries.", len(entries)))
	return entries, nil
}
