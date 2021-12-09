package controllers

import (
	"errors"

	"github.com/EvisuXiao/andrews-common/controller"
	"github.com/EvisuXiao/andrews-common/utils"
	"github.com/gin-gonic/gin"

	"andrews-static/services"
)

type Upload struct {
	controller.Base
}

var (
	uploadController = &Upload{}
)

func NewUploadController() *Upload {
	return uploadController
}

func (c *Upload) Upload(ctx *gin.Context) bool {
	form, err := ctx.MultipartForm()
	if utils.HasErr(err) {
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
