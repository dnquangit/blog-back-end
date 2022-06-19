package postservice

import (
	"go-module/component"
)

func (service *postService) DeletePost(id string) error {

	if post, _ := service.postRepo.FindPostById(id); post == nil {
		return component.NewAppError("post's not found", component.ErrorEntityNotFound.String(), "")
	}

	return service.postRepo.SoftDeletePost(id)
}
