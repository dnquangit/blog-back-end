package userrepo

import (
	"context"
	"go-module/component"
	models "go-module/modules/user/model"
	"gorm.io/gorm"
)

func (repo *userRepo) FindUserById(ctx context.Context, id string) (*models.User, error) {
	db := repo.db
	var user models.User

	if err := db.Table(models.User{}.TableName()).Where("id = ? AND deleted = ?", id, false).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, component.NewAppError(component.ErrMessageNotFoundFromDB.String(), component.ErrorDb.String(), err.Error())
		}
		return nil, component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}

	return &user, nil
}
