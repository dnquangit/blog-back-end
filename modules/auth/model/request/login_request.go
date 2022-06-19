package authrequestmodel

import "go-module/component"

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (request *LoginRequest) Valid() error {
	if request.UserName == "" {
		return component.NewAppError("username cannot be empty", component.ErrorInvalidPayload.String(), "validate payload fail")
	}
	if request.Password == "" {
		return component.NewAppError("password cannot be empty", component.ErrorInvalidPayload.String(), "validate payload fail")
	}

	return nil
}
