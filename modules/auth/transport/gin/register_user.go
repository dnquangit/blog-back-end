package ginauth

import (
	"github.com/gin-gonic/gin"
	"go-module/common"
	"go-module/component"
	models "go-module/modules/auth/model/request"
	service "go-module/modules/auth/service"
	"go-module/modules/user/repo/user"
	"net/http"
)

func RegisterUser(ctx *component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.RegisterUserRequest

		if err := c.ShouldBind(&request); err != nil {
			panic(component.NewAppError(component.ErrMessageInvalidPayload.String(), component.ErrorInvalidPayload.String(), err.Error()))
		}

		userRepo := userrepo.NewUserRepo(ctx)
		service := service.NewAuthService(
			service.WithUserRepo(userRepo),
			service.WithContext(c.Request.Context()),
		)

		if err := service.RegisterUser(&request); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessEmtyDataResponse("register success", http.StatusOK))
	}
}
