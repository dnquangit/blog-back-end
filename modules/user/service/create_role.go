package userservice

import (
	"github.com/google/uuid"
	"go-module/component"
	models "go-module/modules/user/model"
	requestmodels "go-module/modules/user/model/request"
)

func (s *userService) CreateRole(request *requestmodels.CreateRoleRequest) (*requestmodels.CreateRoleResponse, error) {
	if err := request.Valid(); err != nil {
		return nil, err
	}

	role, _ := s.roleRepo.FindRole(s.ctx, request.Name)
	if role != nil {
		return nil, component.NewAppError("role exits", component.ErrorEntityExists.String(), "")
	}

	var roleCreate = models.Role{}
	roleCreate.Id = uuid.NewString()
	roleCreate.Name = request.Name
	if err := s.roleRepo.CreateRole(s.ctx, &roleCreate); err != nil {
		return nil, err
	}

	return &requestmodels.CreateRoleResponse{Id: roleCreate.Id}, nil
}
