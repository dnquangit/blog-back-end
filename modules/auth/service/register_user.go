package authservice

import (
	"github.com/devfeel/mapper"
	"github.com/google/uuid"
	"go-module/common"
	"go-module/component"
	models "go-module/modules/auth/model/request"
	usermodels "go-module/modules/user/model"
)

func (s *authService) RegisterUser(request *models.RegisterUserRequest) error {
	if err := request.Valid(); err != nil {
		return err
	}

	userSameEmails, _ := s.userRepo.FindUsersWithCondition(s.ctx, &usermodels.FilterUserNameEmail{UserName: request.UserName, Email: request.Email})
	if userSameEmails != nil {
		return component.NewAppError("post exits with username or email", component.ErrorEntityExists.String(), "")
	}

	user := usermodels.User{}
	user.Id = uuid.NewString()
	mapper.Mapper(request, &user)
	user.Password, _ = common.HashString(request.Password)

	return s.userRepo.CreateUser(s.ctx, &user)
}
