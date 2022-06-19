package postservice

import (
	"github.com/google/uuid"
	"go-module/common"
	"go-module/component"
	models "go-module/modules/post/model"
	requestmodels "go-module/modules/post/model/request"
	"time"
)

func (service *postService) CreatePost(request *requestmodels.CreatePostRequest) (*requestmodels.CreatePostResponse, error) {
	if err := request.Valid(); err != nil {
		return nil, err
	}

	filterTitle := models.FilterTitle{Title: request.Title}
	existingPosts, _ := service.postRepo.FindPostsWithCondition(&filterTitle)
	if existingPosts != nil && len(*existingPosts) != 0 {
		return nil, component.NewAppError("post exists with title", component.ErrorEntityExists.String(), "")
	}

	post := models.Post{}
	common.Mapper(request, &post)

	post.Id = uuid.New().String()
	if post.Published {
		post.PublishedAt = common.ToPtr(time.Now())
	}

	if err := service.postRepo.CreatePost(&post); err != nil {
		return nil, err
	}

	return &requestmodels.CreatePostResponse{Id: post.Id}, nil
}
