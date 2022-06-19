package postrequestmodel

import "go-module/component"

type UpdatePostRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
}

func (request *UpdatePostRequest) Valid() error {
	if request.Title == "" {
		return component.NewAppError("title cannot be empty", component.ErrorInvalidPayload.String(), "")
	}
	if request.Content == "" {
		return component.NewAppError("content cannot be empty", component.ErrorInvalidPayload.String(), "")
	}
	return nil
}
