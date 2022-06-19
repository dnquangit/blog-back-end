package postservice

import (
	"go-module/common"
	models "go-module/modules/post/model/response"
)

func (service *postService) GetPost(id string) (*models.PostResponse, error) {
	var response models.PostResponse

	post, err := service.postRepo.FindPostById(id)
	if err != nil {
		return nil, err
	}

	common.Mapper(post, &response)

	return &response, nil
}
