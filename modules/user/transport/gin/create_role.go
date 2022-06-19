package ginuser

import (
	"github.com/gin-gonic/gin"
	"go-module/common"
	"go-module/component"
	requestmodels "go-module/modules/user/model/request"
	"go-module/modules/user/repo/role"
	service "go-module/modules/user/service"
	"net/http"
)

func CreateRole(ctx *component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request requestmodels.CreateRoleRequest

		if err := c.ShouldBind(&request); err != nil {
			panic(component.NewAppError(component.ErrMessageInvalidPayload.String(), component.ErrorInvalidPayload.String(), err.Error()))
		}

		service := service.NewUserService(
			service.WithRoleRepo(rolerepo.NewRoleRepo(ctx)),
			service.WithContext(c.Request.Context()),
		)

		role, err := service.CreateRole(&request)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewSuccessResponse(role, "create role success", http.StatusCreated))
	}
}
