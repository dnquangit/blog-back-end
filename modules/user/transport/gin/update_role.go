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

func UpdateRole(ctx *component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request requestmodels.UpdateRoleRequest

		if err := c.ShouldBind(&request); err != nil {
			panic(component.NewAppError(component.ErrMessageInvalidPayload.String(), component.ErrorInvalidPayload.String(), err.Error()))
		}

		id := c.Param("id")

		service := service.NewUserService(service.WithRoleRepo(rolerepo.NewRoleRepo(ctx)))
		if err := service.UpdateRole(id, &request); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.NewSuccessEmtyDataResponse("update role success", http.StatusOK))
	}
}
