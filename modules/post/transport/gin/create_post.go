package ginpost

import (
	"github.com/gin-gonic/gin"
	"go-module/common"
	"go-module/component"
	models "go-module/modules/post/model/request"
	repo "go-module/modules/post/repo"
	postservice "go-module/modules/post/service"
	"net/http"
)

func CreatePost(ctx *component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.CreatePostRequest

		if err := c.ShouldBind(&request); err != nil {
			panic(component.NewAppError(component.ErrMessageInvalidPayload.String(), component.ErrorInvalidPayload.String(), err.Error()))
		}

		postRepo := repo.NewPostRepo(ctx)
		service := postservice.NewPostService(
			postservice.WithPostRepo(postRepo),
			postservice.WithContext(c.Request.Context()),
		)
		post, err := service.CreatePost(&request)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewSuccessResponse(post, "create success", http.StatusCreated))
	}
}
