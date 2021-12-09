package controllers

import (
	"github.com/EvisuXiao/andrews-common/controller"
	"github.com/EvisuXiao/andrews-common/utils"
	"github.com/gin-gonic/gin"

	"andrews-static/services"
	"andrews-static/types"
)

type Preview struct {
	controller.Base
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
	if utils.HasErr(err) {
		return c.InvalidParamResponse(ctx, err)
	}
	previewer := services.NewPreviewer(uriForm.Filename, uriForm.MediaType)
	return c.StaticFile(ctx, previewer.Preview())
}
