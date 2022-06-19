package ginpost

import (
	"github.com/gin-gonic/gin"
	"go-module/common"
	"go-module/component"
	models "go-module/modules/post/model"
	responsemodels "go-module/modules/post/model/response"
	repo "go-module/modules/post/repo"
	postservice "go-module/modules/post/service"
	"net/http"
	"strconv"
)

func ListPost(ctx *component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		publishedParam := c.Query("published")
		published, _ := strconv.ParseBool(publishedParam)

		postRepo := repo.NewPostRepo(ctx)
		service := postservice.NewPostService(
			postservice.WithPostRepo(postRepo),
			postservice.WithContext(c.Request.Context()),
		)

		var posts *[]responsemodels.PostResponse
		var err error
		if published {
			posts, err = service.ListPost(&models.FilterPublished{Published: published})
		} else {
			posts, err = service.ListPost(&models.FilterAll{})
		}

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(posts, "get success", http.StatusOK))
	}
}
