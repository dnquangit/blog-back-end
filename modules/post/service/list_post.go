package postservice

import (
	"go-module/common"
	models "go-module/modules/post/model/response"
)

func (service *postService) ListPost(filter interface{}) (*[]models.PostResponse, error) {

	posts, err := service.postRepo.FindPostsWithCondition(filter)
	if err != nil {
		return nil, err
	}

	response := make([]models.PostResponse, len(*posts))
	common.MapperSlice(*posts, response)

	return &response, nil
}
