package controllers

import (
	"github.com/gin-gonic/gin"
	"upload-test/pkg/utils"
	"upload-test/services"
	"upload-test/types"
)

type Preview struct {
	Base
}

var (
	previewController = &Preview{}
)

func NewPreviewController() *Preview {
	return previewController
}

func (c *Preview) Preview(ctx *gin.Context) bool {
	var uriForm struct {
		MediaType types.MEDIA_TYPE `uri:"mediaType" binding:"required"`
		Filename  string           `uri:"filename" binding:"required"`
	}
	err := ctx.ShouldBindUri(&uriForm)
	if !utils.IsEmpty(err) {
		return c.InvalidParamResponse(ctx, err)
	}
	previewer := services.NewPreviewer(uriForm.Filename, uriForm.MediaType)
	return c.StaticFile(ctx, previewer.Preview())
}
