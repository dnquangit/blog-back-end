package ginuser

import (
	"github.com/gin-gonic/gin"
	"go-module/common"
	"go-module/component"
	requestmodels "go-module/modules/user/model/request"
	"go-module/modules/user/repo/role"
	"go-module/modules/user/repo/user"
	service "go-module/modules/user/service"
	"net/http"
)

func CreateRoleForUser(ctx *component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request requestmodels.CreateRoleUserRequest

		if err := c.ShouldBind(&request); err != nil {
			panic(component.NewAppError(component.ErrMessageInvalidPayload.String(), component.ErrorInvalidPayload.String(), err.Error()))
		}

		id := common.GetClaims(c)["id"]

		service := service.NewUserService(
			service.WithRoleRepo(rolerepo.NewRoleRepo(ctx)),
			service.WithUserRepo(userrepo.NewUserRepo(ctx)),
		)

		if err := service.CreateRoleForUser(id, &request); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.NewSuccessEmtyDataResponse("create role for user success", http.StatusOK))
	}
}
