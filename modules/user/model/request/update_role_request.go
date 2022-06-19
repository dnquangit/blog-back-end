package userrequestmodel

import "go-module/component"

type UpdateRoleRequest struct {
	Name string `json:"name"`
}

func (request *UpdateRoleRequest) Valid() error {
	if request.Name == "" {
		return component.NewAppError("name cannot be empty", component.ErrorInvalidPayload.String(), "")
	}

	return nil
}
