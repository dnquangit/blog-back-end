package rolerepo

import (
	"context"
	"go-module/component"
	models "go-module/modules/user/model"
)

func (repo *roleRepo) CreateRole(ctx context.Context, role *models.Role) error {
	db := repo.db

	if err := db.Table(models.Role{}.TableName()).Select("id", "name").Create(&role).Error; err != nil {
		return component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}
	return nil
}
