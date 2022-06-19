package ginuser

import (
	"github.com/gin-gonic/gin"
	"go-module/common"
	"go-module/component"
	"go-module/modules/user/repo/role"
	service "go-module/modules/user/service"
	"net/http"
)

func GetRoles(ctx *component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		claims := common.GetClaims(c)
		userId := claims["id"]

		service := service.NewUserService(service.WithRoleRepo(rolerepo.NewRoleRepo(ctx)))

		roles, err := service.GetRoles(userId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(roles, "get roles success", http.StatusOK))
	}
}
