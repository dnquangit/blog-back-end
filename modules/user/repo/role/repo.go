package rolerepo

import (
	"go-module/component"
	"gorm.io/gorm"
)

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(ctx *component.AppContext) *roleRepo {
	return &roleRepo{db: ctx.GormDB}
}
