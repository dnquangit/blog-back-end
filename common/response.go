package common

import (
	"go-module/component"
	"net/http"
)

type response struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	Reason     string      `json:"reason"`
}

func NewSuccessResponse(data interface{}, message string, code int) *response {
	return &response{Data: data, Message: message, StatusCode: code, Status: "success"}
}

func NewSuccessEmtyDataResponse(message string, code int) *response {
	return &response{Data: nil, Message: message, StatusCode: code, Status: "success"}
}

func NewBadRequestResponse(appError *component.AppError) *response {
	return &response{Data: nil, Message: appError.Message, Reason: appError.Reason, StatusCode: http.StatusBadRequest, Status: "fail"}
}

func NewNotFoundResponse(appError *component.AppError) *response {
	return &response{Data: nil, Message: appError.Message, Reason: appError.Reason, StatusCode: http.StatusNotFound, Status: "fail"}
}

func NewInternalErrorResponse(err error) *response {
	return &response{Data: nil, Message: "There are some internal error", Reason: err.Error(), StatusCode: http.StatusInternalServerError, Status: "fail"}
}

func NewForbiddenResponse(appError *component.AppError) *response {
	return &response{Data: nil, Message: appError.Message, Reason: appError.Reason, StatusCode: http.StatusForbidden, Status: "fail"}
}

func NewUnAuthorizedResponse(appError *component.AppError) *response {
	return &response{Data: nil, Message: appError.Message, Reason: appError.Reason, StatusCode: http.StatusUnauthorized, Status: "fail"}
}
