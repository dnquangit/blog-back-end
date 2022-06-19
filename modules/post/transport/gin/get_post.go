package ginpost

import (
	"github.com/gin-gonic/gin"
	"go-module/common"
	"go-module/component"
	repo "go-module/modules/post/repo"
	postservice "go-module/modules/post/service"
	"net/http"
)

func GetPost(ctx *component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		postRepo := repo.NewPostRepo(ctx)
		service := postservice.NewPostService(
			postservice.WithPostRepo(postRepo),
			postservice.WithContext(c.Request.Context()),
		)
		post, err := service.GetPost(id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(post, "get success", http.StatusOK))
	}
}
