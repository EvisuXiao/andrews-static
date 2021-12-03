package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"upload-test/config"
	"upload-test/pkg/logging"
	"upload-test/pkg/utils"
)

// 为了每次api返回错误json时不用另起一行加return, 将默认handler添加返回值
type handlerFunc func(*gin.Context) bool

type routerConf struct {
	method   string
	path     string
	handlers interface{} // handlerFunc, []handlerFunc
}

type routerGroup struct {
	path       string
	middleware interface{} // handlerFunc, []handlerFunc
	routers    []*routerConf
}

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	setMode()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello Andrews")
		ctx.Abort()
	})
	initRouter(r.Group(""), initWeb())
	return r
}

func setMode() {
	if config.IsLocalEnv() {
		gin.SetMode(gin.DebugMode)
	} else if config.IsProdEnv() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.TestMode)
	}
}

// 将自定义handler转回默认handler
func routersHandler(handlers interface{}) []gin.HandlerFunc {
	rawFunc := make([]gin.HandlerFunc, 0)
	if utils.IsEmpty(handlers) {
		return rawFunc
	}
	var customFunc []handlerFunc
	switch v := handlers.(type) {
	case func(*gin.Context) bool:
		customFunc = []handlerFunc{v}
	case []handlerFunc:
		customFunc = v
	default:
		logging.Fatal("Init router fatal: unsupported handler type")
	}
	for _, f := range customFunc {
		rawFunc = append(rawFunc, routerHandler(f))
	}
	return rawFunc
}

func routerHandler(handler handlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c)
	}
}

func initRouter(engine *gin.RouterGroup, routers []*routerGroup) {
	for _, group := range routers {
		apiGroup := engine.Group(group.path, routersHandler(group.middleware)...)
		for _, router := range group.routers {
			ginHandlers := routersHandler(router.handlers)
			switch router.method {
			case http.MethodGet:
				apiGroup.GET(router.path, ginHandlers...)
			case http.MethodPost:
				apiGroup.POST(router.path, ginHandlers...)
			case http.MethodPut:
				apiGroup.PUT(router.path, ginHandlers...)
			case http.MethodDelete:
				apiGroup.DELETE(router.path, ginHandlers...)
			default:
				apiGroup.Any(router.path, ginHandlers...)
			}
		}
	}
}
