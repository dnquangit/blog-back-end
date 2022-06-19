package rolerepo

import (
	"context"
	"go-module/component"
	models "go-module/modules/user/model"
)

func (repo *roleRepo) CreateUserRole(ctx context.Context, userRole *models.UserRole) error {
	db := repo.db

	if err := db.Table(models.UserRole{}.TableName()).Select("user_id", "role_id").Create(&userRole).Error; err != nil {
		return component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}
	return nil
}
