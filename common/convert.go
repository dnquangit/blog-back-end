package common

import (
	models "go-module/modules/user/model"
	"strings"
)

func ToPtr[T any](t T) *T {
	return &t
}

func GetRoles(roles *[]models.Role) string {
	if roles == nil {
		return ""
	}

	result := ""

	for _, role := range *roles {
		result = result + "," + role.Name
	}

	return strings.Trim(result, ",")
}

func IsError(input interface{}) bool {
	value, _ := input.(error)
	return value != nil
}

func ToError(input interface{}) error {
	value, _ := input.(error)
	return value
}
