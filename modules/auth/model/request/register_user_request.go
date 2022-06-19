package authrequestmodel

import (
	"go-module/component"
)

type RegisterUserRequest struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (request *RegisterUserRequest) Valid() error {
	if request.UserName == "" {
		return component.NewAppError("username cannot be empty", component.ErrorInvalidPayload.String(), "validate payload fail")
	}
	if request.FirstName == "" {
		return component.NewAppError("first name cannot be empty", component.ErrorInvalidPayload.String(), "validate payload fail")
	}
	if request.LastName == "" {
		return component.NewAppError("last name cannot be empty", component.ErrorInvalidPayload.String(), "validate payload fail")
	}
	if request.Email == "" {
		return component.NewAppError("email cannot be empty", component.ErrorInvalidPayload.String(), "validate payload fail")
	}
	return nil
}
