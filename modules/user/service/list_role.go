package userservice

import (
	"go-module/common"
	responsemodels "go-module/modules/user/model/response"
)

func (s *userService) GetRoles(userId string) (*[]responsemodels.RoleResponse, error) {

	roles, err := s.roleRepo.FindRolesByUserId(s.ctx, userId)
	if err != nil {
		return nil, err
	}

	responses := make([]responsemodels.RoleResponse, len(*roles))
	common.MapperSlice(*roles, responses)

	return &responses, nil
}
