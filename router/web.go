package router

import (
	"net/http"
	"upload-test/controllers"
)

func initWeb() []*routerGroup {
	previewController := controllers.NewPreviewController()
	uploadController := controllers.NewUploadController()
	return []*routerGroup{
		{
			"",
			[]handlerFunc{},
			[]*routerConf{
				{http.MethodGet, ":mediaType/:filename", previewController.Preview},
			},
		},
		{
			"upload",
			[]handlerFunc{},
			[]*routerConf{
				{http.MethodPost, "", uploadController.Upload},
			},
		},
	}
}
