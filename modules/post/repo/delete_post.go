package postrepo

import (
	"go-module/component"
	models "go-module/modules/post/model"
)

func (repo *postRepo) SoftDeletePost(id string) error {
	db := repo.db
	if err := db.Table(models.Post{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"deleted": true}).Error; err != nil {
		return component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}
	return nil
}
