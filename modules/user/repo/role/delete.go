package rolerepo

import (
	"context"
	"go-module/component"
	models "go-module/modules/user/model"
)

func (repo *roleRepo) SoftDeleteRole(ctx context.Context, id string) error {
	db := repo.db
	if err := db.Table(models.Role{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"deleted": true}).Error; err != nil {
		return component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}
	return nil
}
