package middleware

import (
	_ "errors"
	"github.com/gin-gonic/gin"
	"go-module/common"
	"go-module/component"
	"net/http"
)

func Recover(context *component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil && common.IsError(err) {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*component.AppError); ok {
					writeAppError(c, appErr)
					return
				}

				c.AbortWithStatusJSON(http.StatusInternalServerError, common.NewInternalErrorResponse(common.ToError(err)))
			}
		}()

		c.Next()
	}
}

func writeAppError(c *gin.Context, err *component.AppError) {
	switch typeMessage := err.Type; typeMessage {
	case component.ErrorInvalidPayload.String(), component.ErrorInvalidAuth.String(), component.ErrorEntityExists.String():
		c.AbortWithStatusJSON(http.StatusBadRequest, common.NewBadRequestResponse(err))
	case component.ErrorEntityNotFound.String():
		c.AbortWithStatusJSON(http.StatusNotFound, common.NewNotFoundResponse(err))
	case component.ErrNoPermission.String():
		c.AbortWithStatusJSON(http.StatusForbidden, common.NewForbiddenResponse(err))
	}
}
