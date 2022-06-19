package common

import (
	"github.com/gin-gonic/gin"
	"go-module/component/tokenprovider/jwt"
	"strings"
)

func GetClaims(c *gin.Context) map[string]string {

	var token string
	if c.Request.Header["Authorization"] != nil {
		token = c.Request.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", -1)
	}

	if token == "" {
		return nil
	}

	return jwt.GetClaimsFromToken(token)
}
