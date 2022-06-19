package userrequestmodel

import "go-module/component"

type CreateRoleUserRequest struct {
	Name string `json:"name"`
}

func (request *CreateRoleUserRequest) Valid() error {
	if request.Name == "" {
		return component.NewAppError("name cannot be empty", component.ErrorInvalidPayload.String(), "")
	}

	return nil
}
