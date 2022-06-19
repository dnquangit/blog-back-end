package userservice

import (
	"go-module/component"
	requestmodels "go-module/modules/user/model/request"
)

func (s *userService) UpdateRole(id string, request *requestmodels.UpdateRoleRequest) error {
	if err := request.Valid(); err != nil {
		return nil
	}

	role, _ := s.roleRepo.FindRoleById(s.ctx, id)
	if role == nil {
		return component.NewAppError("role isn't found", component.ErrorEntityNotFound.String(), "")
	}

	roleWithName, _ := s.roleRepo.FindRole(s.ctx, request.Name)
	if roleWithName != nil {
		return component.NewAppError("role exits", component.ErrorEntityExists.String(), "")
	}

	role.Name = request.Name

	return s.roleRepo.UpdateRole(s.ctx, role)
}
