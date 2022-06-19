package postservice

import (
	"fmt"
	"go-module/common"
	"go-module/component"
	models "go-module/modules/post/model"
	requestmodels "go-module/modules/post/model/request"
	"time"
)

func (s *postService) UpdatePost(id string, request *requestmodels.UpdatePostRequest) error {
	if err := request.Valid(); err != nil {
		return err
	}

	updatePublishedDate := false
	existingPost, _ := s.postRepo.FindPostById(id)
	fmt.Println("existingPost")
	fmt.Println(existingPost.Published)

	if existingPost == nil {
		return component.NewAppError("post's not found", component.ErrorEntityNotFound.String(), "")
	}

	postSameTitle, _ := s.postRepo.FindPostsWithCondition(&models.FilterTitle{Title: request.Title, Id: id})
	if postSameTitle != nil && len(*postSameTitle) > 0 {
		return component.NewAppError("post exits with title", component.ErrorEntityExists.String(), "")
	}

	fmt.Println("request")
	fmt.Println(request.Published)

	fmt.Println("request.Published != existingPost.Published")
	fmt.Println(request.Published != existingPost.Published)
	updatePublishedDate = request.Published != existingPost.Published
	common.Mapper(request, existingPost)

	if updatePublishedDate {
		if request.Published {
			existingPost.PublishedAt = common.ToPtr(time.Now())
		} else {
			existingPost.Published = false
			existingPost.PublishedAt = nil
		}
	}

	return s.postRepo.UpdatePost(existingPost)
}
