package userrequestmodel

import "go-module/component"

type CreateRoleRequest struct {
	Name string `json:"name"`
}

type CreateRoleResponse struct {
	Id string `json:"id"`
}

func (request *CreateRoleRequest) Valid() error {
	if request.Name == "" {
		return component.NewAppError("name cannot be empty", component.ErrorInvalidPayload.String(), "")
	}

	return nil
}
