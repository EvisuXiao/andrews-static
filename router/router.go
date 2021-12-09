package router

import (
	"net/http"

	"github.com/EvisuXiao/andrews-common/router"

	"andrews-static/controllers"
)

func Init() *router.MainGroup {
	previewController := controllers.NewPreviewController()
	uploadController := controllers.NewUploadController()
	groups := []*router.Group{
		{
			"",
			[]router.Handler{},
			[]*router.Item{
				{http.MethodGet, ":mediaType/:filename", previewController.Preview},
			},
		},
		{
			"upload",
			[]router.Handler{},
			[]*router.Item{
				{http.MethodPost, "", uploadController.Upload},
			},
		},
	}
	return &router.MainGroup{Groups: groups}
}
