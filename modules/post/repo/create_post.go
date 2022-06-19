package postrepo

import (
	"go-module/component"
	models "go-module/modules/post/model"
)

func (repo *postRepo) CreatePost(post *models.Post) error {
	db := repo.db

	if err := db.Table(models.Post{}.TableName()).Select("id", "title", "content", "published", "published_at").Create(&post).Error; err != nil {
		return component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}
	return nil
}
