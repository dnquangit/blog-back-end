package userrepo

import (
	"context"
	"go-module/component"
	models "go-module/modules/user/model"
)

func (repo *userRepo) CreateUser(ctx context.Context, user *models.User) error {
	db := repo.db

	if err := db.Table(models.User{}.TableName()).Select("id", "username", "password", "email", "first_name", "last_name").Create(&user).Error; err != nil {
		return component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}
	return nil
}
