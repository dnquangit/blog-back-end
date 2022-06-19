package rolerepo

import (
	"context"
	"go-module/component"
	models "go-module/modules/user/model"
)

func (repo *roleRepo) UpdateRole(ctx context.Context, role *models.Role) error {
	db := repo.db
	if err := db.Table(models.Role{}.TableName()).Save(role).Error; err != nil {
		return component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}
	return nil
}
