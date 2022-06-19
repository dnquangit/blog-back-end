package postrepo

import (
	"fmt"
	"go-module/component"
	models "go-module/modules/post/model"
)

func (repo *postRepo) UpdatePost(post *models.Post) error {
	fmt.Println(post)
	db := repo.db

	if err := db.Table(models.Post{}.TableName()).Save(post).Error; err != nil {
		return component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}

	return nil
}
