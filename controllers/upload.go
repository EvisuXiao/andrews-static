package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"upload-test/pkg/utils"
	"upload-test/services"
)

type Upload struct {
	Base
}

var (
	uploadController = &Upload{}
)

func NewUploadController() *Upload {
	return uploadController
}

func (c *Upload) Upload(ctx *gin.Context) bool {
	form, err := ctx.MultipartForm()
	if !utils.IsEmpty(err) {
		return c.InvalidParamResponse(ctx, err)
	}
	files, ok := form.File["file"]
	if !ok {
		return c.InvalidParamResponse(ctx, errors.New("can't get files with file key"))
	}
	for _, file := range files {
		uploader := services.NewUploader(file)
		uploader.RegisterFilter("thumb", services.ThumbFilter(100, 100))
		err := uploader.Upload()
		if !utils.IsEmpty(err) {
			return c.FailureResponse(ctx, err)
		}
	}
	return c.SuccessResponse(ctx)
}
