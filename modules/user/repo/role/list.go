package rolerepo

import (
	"context"
	"go-module/component"
	models "go-module/modules/user/model"
	"gorm.io/gorm"
)

func (repo *roleRepo) FindRolesByUserId(ctx context.Context, userId string) (*[]models.Role, error) {
	var roles []models.Role
	db := repo.db
	if err := db.Table(models.UserRole{}.TableName()).Select("roles.name, roles.id, roles.created_at, roles.updated_at, roles.deleted").Joins("left join roles on roles.id = user_roles.role_id").Where("user_id = ?", userId).Find(&roles).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, component.NewAppError(component.ErrMessageNotFoundFromDB.String(), component.ErrorDb.String(), err.Error())
		}
		return nil, component.NewAppError(component.ErrMessageFromDB.String(), component.ErrorDb.String(), err.Error())
	}

	return &roles, nil
}
