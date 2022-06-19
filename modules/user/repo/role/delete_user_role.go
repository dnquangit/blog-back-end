package rolerepo

import (
	"context"
	"go-module/component"
	models "go-module/modules/user/model"
)

func (repo *roleRepo) SoftDeleteUserRoles(ctx context.Context, userRoles *models.UserRole) error {
	db := repo.db

	if err := db.Table(models.UserRole{}.TableName()).Where("user_id = ? AND role_id = ?", userRoles.UserId, userRoles.RoleId).Updates(map[string]interface{}{"deleted": true}).Error; err != nil {
		return component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}
	return nil
}
