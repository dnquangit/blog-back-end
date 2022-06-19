package postrequestmodel

import "go-module/component"

type CreatePostRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
}

type CreatePostResponse struct {
	Id string `json:"id"`
}

func (request *CreatePostRequest) Valid() error {
	if request.Title == "" {
		return component.NewAppError("title cannot be empty", component.ErrorInvalidPayload.String(), "")
	}
	if request.Content == "" {
		return component.NewAppError("content cannot be empty", component.ErrorInvalidPayload.String(), "")
	}
	return nil
}
