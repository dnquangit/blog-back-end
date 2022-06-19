package postservice

import (
	"context"
	models "go-module/modules/post/model"
)

type PostRepo interface {
	FindPostById(string) (*models.Post, error)
	FindPostsWithCondition(interface{}) (*[]models.Post, error)
	CreatePost(*models.Post) error
	SoftDeletePost(string) error
	UpdatePost(*models.Post) error
}

type postService struct {
	postRepo PostRepo
	ctx      context.Context
}

type Configurations func(*postService)

func NewPostService(options ...Configurations) *postService {
	u := &postService{}
	for _, option := range options {
		option(u)
	}
	return u
}

func WithPostRepo(repo PostRepo) Configurations {
	return func(service *postService) {
		service.postRepo = repo
	}
}

func WithContext(context context.Context) Configurations {
	return func(service *postService) {
		service.ctx = context
	}
}
