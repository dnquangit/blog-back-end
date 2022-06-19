package userservice

import (
	"go-module/component"
	models "go-module/modules/user/model"
	requestmodels "go-module/modules/user/model/request"
)

func (s *userService) CreateRoleForUser(userId string, request *requestmodels.CreateRoleUserRequest) error {
	if err := request.Valid(); err != nil {
		return nil
	}

	user, _ := s.userRepo.FindUserById(s.ctx, userId)
	if user == nil {
		return component.NewAppError("user isn't found", component.ErrorEntityNotFound.String(), "")
	}

	roleWithName, _ := s.roleRepo.FindRole(s.ctx, request.Name)
	if roleWithName == nil {
		return component.NewAppError("role isn't found", component.ErrorEntityNotFound.String(), "")
	}

	userRole := models.UserRole{UserId: userId, RoleId: roleWithName.Id}

	return s.roleRepo.CreateUserRole(s.ctx, &userRole)
}
