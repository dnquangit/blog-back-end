package postrepo

import (
	"context"
	"go-module/component"
	"gorm.io/gorm"
)

type postRepo struct {
	db  *gorm.DB
	ctx context.Context
}

func NewPostRepo(ctx *component.AppContext) *postRepo {
	return &postRepo{db: ctx.GormDB}
}

type Configurations func(*postRepo)

func WithContext(context context.Context) Configurations {
	return func(service *postRepo) {
		service.ctx = context
	}
}
