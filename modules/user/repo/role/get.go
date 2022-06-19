package rolerepo

import (
	"context"
	"go-module/component"
	models "go-module/modules/user/model"
	"gorm.io/gorm"
)

func (repo *roleRepo) FindRole(ctx context.Context, roleName string) (*models.Role, error) {
	db := repo.db

	var role models.Role
	if err := db.Table(models.Role{}.TableName()).Where("name = ? AND deleted = ?", roleName, false).First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, component.NewAppError(component.ErrMessageNotFoundFromDB.String(), component.ErrorEntityNotFound.String(), err.Error())
		}
		return nil, component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}
	return &role, nil
}

func (repo *roleRepo) FindRoleById(ctx context.Context, id string) (*models.Role, error) {
	db := repo.db

	var role models.Role
	if err := db.Table(models.Role{}.TableName()).Where("id = ? AND deleted = ?", id, false).First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, component.NewAppError(component.ErrMessageNotFoundFromDB.String(), component.ErrorEntityNotFound.String(), err.Error())
		}
		return nil, component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}
	return &role, nil
}
