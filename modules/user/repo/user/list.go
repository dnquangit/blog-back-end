package userrepo

import (
	"context"
	"go-module/component"
	models "go-module/modules/user/model"
)

func (repo *userRepo) FindUsersWithCondition(ctx context.Context, filter interface{}) (*[]models.User, error) {

	switch filter.(type) {
	case *models.FilterUserNameEmail:
		return repo.FindUserByFilterUserNameEmail(ctx, models.NewFilterUserNameEmail(filter))
	case *models.FilterUserName:
		return repo.FindUserByFilterUserName(ctx, models.NewFilterUserName(filter))
	}

	return nil, nil
}

func (repo *userRepo) FindUserByFilterUserNameEmail(ctx context.Context, filter *models.FilterUserNameEmail) (*[]models.User, error) {
	if filter != nil {
		var users []models.User
		db := repo.db
		if err := db.Table(models.User{}.TableName()).Where("username = ? OR email = ? AND deleted = ?", filter.UserName, filter.Email, false).Find(&users).Error; err != nil {
			return &users, component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
		}
	}

	return nil, nil
}

func (repo *userRepo) FindUserByFilterUserName(ctx context.Context, filter *models.FilterUserName) (*[]models.User, error) {
	if filter != nil {
		var users []models.User
		db := repo.db
		if err := db.Table(models.User{}.TableName()).Where("username = ? AND deleted = ?", filter.UserName, false).Find(&users).Error; err != nil {
			return nil, component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
		}
		return &users, nil
	}

	return nil, nil
}
