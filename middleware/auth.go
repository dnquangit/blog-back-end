package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-module/common"
	"go-module/component"
	repo "go-module/modules/user/repo/user"
)

func RequiredAuth(ctx *component.AppContext, role string) gin.HandlerFunc {
	return func(c *gin.Context) {

		claims := common.GetClaims(c)

		if claims["role"] != role || !isValidUser(claims, c.Request.Context(), ctx) {
			panic(component.NewAppError("user isn't allowed", component.ErrNoPermission.String(), ""))
		}

		c.Next()
	}
}

func isValidUser(claims map[string]string, context context.Context, ctx *component.AppContext) bool {
	userId := claims["id"]
	fmt.Println("userId")
	fmt.Println(userId)
	repo := repo.NewUserRepo(ctx)
	user, _ := repo.FindUserById(context, userId)
	return user != nil
}
