package postrepo

import (
	"go-module/component"
	models "go-module/modules/post/model"
	"gorm.io/gorm"
)

func (repo *postRepo) FindPostById(id string) (*models.Post, error) {
	db := repo.db
	var post models.Post

	if err := db.Table(models.Post{}.TableName()).Where("id = ? AND deleted = ?", id, false).First(&post).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, component.NewAppError(component.ErrMessageNotFoundFromDB.String(), component.ErrorEntityNotFound.String(), err.Error())
		}
		return nil, component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}

	return &post, nil
}
