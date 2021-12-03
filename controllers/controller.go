package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"upload-test/pkg/exception"
	"upload-test/pkg/translation"
	"upload-test/pkg/utils"
)

type Base struct{}

type ApiOutput struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *Base) SuccessResponse(ctx *gin.Context, data ...interface{}) bool {
	output := NewOutput(exception.SUCCESS_CODE)
	if len(data) > 1 {
		output.SetMessage(fmt.Sprint(data[1]))
	} else {
		output.SetMessage(exception.SUCCESS_MSG)
	}
	if !utils.IsEmpty(data) {
		output.SetData(data[0])
	}
	return output.ApiResponse(ctx)
}

func (c *Base) FailureResponseWithCode(ctx *gin.Context, code int, desc ...interface{}) bool {
	output := NewOutput(code)
	descLen := len(desc)
	if descLen > 0 {
		msg := desc[0]
		if utils.IsEmpty(msg) {
			output.SetMessage(exception.FAILURE_MSG)
		} else if e, ok := msg.(*exception.CustomError); ok {
			output.SetMessage(e.Error())
		} else {
			output.SetMessage(exception.SERVER_ERROR_MSG)
			output.SetData(fmt.Sprint(msg))
		}
		if descLen > 1 && !utils.IsEmpty(desc[1]) {
			if err, ok := desc[1].(validator.ValidationErrors); ok {
				desc[1] = translation.TranslateByEn(err)
			}
			output.SetData(fmt.Sprint(desc[1]))
		}
	} else {
		output.SetMessage(exception.FAILURE_MSG)
	}
	return output.ApiResponse(ctx)
}

func (c *Base) FailureResponse(ctx *gin.Context, desc ...interface{}) bool {
	return c.FailureResponseWithCode(ctx, exception.FAILURE_CODE, desc...)
}

func (c *Base) InvalidParamResponse(ctx *gin.Context, err error) bool {
	return c.FailureResponseWithCode(ctx, exception.PARAM_CODE, exception.INVALID_PARAM_ERR, err)
}

func (c *Base) StaticFile(ctx *gin.Context, filename string) bool {
	ctx.File(filename)
	ctx.Abort()
	return true
}

func (c *Base) Next(ctx *gin.Context) bool {
	ctx.Next()
	return true
}

func NewOutput(code int) *ApiOutput {
	o := &ApiOutput{}
	o.Code = code
	return o
}

func (o *ApiOutput) ApiResponse(ctx *gin.Context) bool {
	ctx.AbortWithStatusJSON(http.StatusOK, o)
	return true
}

func (o *ApiOutput) SetCode(code int) {
	o.Code = code
}

func (o *ApiOutput) SetMessage(msg string) {
	o.Message = msg
}

func (o *ApiOutput) SetData(data interface{}) {
	o.Data = data
}

func IsGet(ctx *gin.Context) bool {
	return ctx.Request.Method == http.MethodGet
}

func IsPost(ctx *gin.Context) bool {
	return ctx.Request.Method == http.MethodPost
}

func IsPut(ctx *gin.Context) bool {
	return ctx.Request.Method == http.MethodPut
}

func IsDelete(ctx *gin.Context) bool {
	return ctx.Request.Method == http.MethodDelete
}
