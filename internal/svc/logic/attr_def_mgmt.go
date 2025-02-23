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

func (adm *AttributeDefMgmt) GetAttributeDefs() ([]meta.AttributeDef, error) {
	return adm.repo.GetAll()
}
