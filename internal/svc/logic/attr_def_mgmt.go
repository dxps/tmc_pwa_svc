package logic

import (
	"github.com/dxps/tmc-pwa/internal/shared/model/meta"
	"github.com/dxps/tmc-pwa/internal/svc/repos"
)

type AttributeDefMgmt struct {
	repo *repos.AttributeDefRepo
}

func NewAttributeDefMgmt(repo *repos.AttributeDefRepo) *AttributeDefMgmt {
	return &AttributeDefMgmt{repo}
}

func (m *AttributeDefMgmt) GetAttributeDefs() ([]*meta.AttributeDef, error) {
	return m.repo.GetAll()
}

func (m *AttributeDefMgmt) AddAttributeDef(entry *meta.AttributeDef) error {
	entry.Id = meta.NewId()
	return m.repo.Add(entry)
}
