package authservice

import (
	"context"
	models "go-module/modules/user/model"
)

type UserRepo interface {
	FindUsersWithCondition(context.Context, interface{}) (*[]models.User, error)
	CreateUser(context.Context, *models.User) error
}

type RoleRepo interface {
	FindRolesByUserId(context.Context, string) (*[]models.Role, error)
}

type authService struct {
	userRepo UserRepo
	roleRepo RoleRepo
	ctx      context.Context
}

type Configurations func(*authService)

func NewAuthService(options ...Configurations) *authService {
	u := &authService{}
	for _, option := range options {
		option(u)
	}
	return u
}

func WithRoleRepo(repo RoleRepo) Configurations {
	return func(service *authService) {
		service.roleRepo = repo
	}
}

func WithUserRepo(repo UserRepo) Configurations {
	return func(service *authService) {
		service.userRepo = repo
	}
}

func WithContext(context context.Context) Configurations {
	return func(service *authService) {
		service.ctx = context
	}
}
