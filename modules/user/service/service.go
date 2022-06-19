package userservice

import (
	"context"
	models "go-module/modules/user/model"
)

type UserRepo interface {
	FindUserById(context.Context, string) (*models.User, error)
}

type RoleRepo interface {
	FindRole(context.Context, string) (*models.Role, error)
	FindRoleById(context.Context, string) (*models.Role, error)
	CreateRole(context.Context, *models.Role) error
	UpdateRole(context.Context, *models.Role) error
	CreateUserRole(context.Context, *models.UserRole) error
	FindRolesByUserId(context.Context, string) (*[]models.Role, error)
}

type userService struct {
	roleRepo RoleRepo
	userRepo UserRepo
	ctx      context.Context
}

func NewUserService(options ...Configurations) *userService {
	u := &userService{}
	for _, option := range options {
		option(u)
	}
	return u
}

type Configurations func(*userService)

func WithRoleRepo(repo RoleRepo) Configurations {
	return func(service *userService) {
		service.roleRepo = repo
	}
}

func WithUserRepo(repo UserRepo) Configurations {
	return func(service *userService) {
		service.userRepo = repo
	}
}

func WithContext(context context.Context) Configurations {
	return func(service *userService) {
		service.ctx = context
	}
}
