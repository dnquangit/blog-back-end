package authservice

import (
	"go-module/common"
	"go-module/component"
	"go-module/component/tokenprovider/jwt"
	models "go-module/modules/auth/model/request"
	usermodels "go-module/modules/user/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Login(request *models.LoginRequest) (string, error) {
	if err := request.Valid(); err != nil {
		return "", err
	}

	users, _ := s.userRepo.FindUsersWithCondition(s.ctx, &usermodels.FilterUserName{UserName: request.UserName})
	if users == nil {
		return "", component.NewAppError("user isn't found", component.ErrorEntityNotFound.String(), "find user fail")
	}

	user := (*users)[0]

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", component.NewAppError("password isn't correct", component.ErrorInvalidAuth.String(), "crypt password fail")
	}

	roles, _ := s.roleRepo.FindRolesByUserId(s.ctx, user.Id)
	token, err := jwt.GenerateJWT(user.Id, common.GetRoles(roles))

	return token, err
}
