package ginauth

import (
	"github.com/gin-gonic/gin"
	"go-module/common"
	"go-module/component"
	models "go-module/modules/auth/model/request"
	service "go-module/modules/auth/service"
	"go-module/modules/user/repo/role"
	"go-module/modules/user/repo/user"
	"net/http"
)

func Login(ctx *component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.LoginRequest

		if err := c.ShouldBind(&request); err != nil {
			panic(component.NewAppError(component.ErrMessageInvalidPayload.String(), component.ErrorInvalidPayload.String(), err.Error()))
		}

		userRepo := userrepo.NewUserRepo(ctx)
		roleRepo := rolerepo.NewRoleRepo(ctx)
		service := service.NewAuthService(
			service.WithUserRepo(userRepo),
			service.WithRoleRepo(roleRepo),
			service.WithContext(c.Request.Context()),
		)

		token, err := service.Login(&request)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(map[string]string{"token": token}, "login success", http.StatusOK))
	}
}
