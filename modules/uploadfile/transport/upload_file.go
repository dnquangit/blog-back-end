package ginupload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-module/common"
	"go-module/component"
	uploadfileresponsemodel "go-module/modules/uploadfile/model/response"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func UploadFile(ctx *component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("UploadFile")
		formFile, err := c.FormFile("file")

		// The file cannot be received.
		if err != nil {
			panic(component.NewAppError(component.ErrMessageInvalidPayload.String(), component.ErrorInvalidPayload.String(), err.Error()))
		}

		// Parse to byte[]
		openedFile, _ := formFile.Open()
		file, err := ioutil.ReadAll(openedFile)
		if err != nil {
			panic(component.NewAppError(component.ErrMessageInvalidPayload.String(), component.ErrorInvalidPayload.String(), err.Error()))
		}

		extension := filepath.Ext(formFile.Filename)
		newFileName := formFile.Filename + uuid.New().String() + extension
		url, err := ctx.UploadFileProvider.UploadFile(c.Request.Context(), newFileName, file)
		if err != nil {
			panic(err)
		}
		response := uploadfileresponsemodel.UploadFileResponse{Url: url}
		c.JSON(http.StatusOK, common.NewSuccessResponse(response, "upload success", http.StatusOK))
	}
}
