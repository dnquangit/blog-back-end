package userrepo

import (
	"go-module/component"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(ctx *component.AppContext) *userRepo {
	return &userRepo{db: ctx.GormDB}
}
