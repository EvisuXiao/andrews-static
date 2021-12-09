package services

import (
	"andrews-static/pkg/image"
	"andrews-static/types"
)

func ThumbFilter(width, height int) UploadFilterHandler {
	return func(file types.FileBrief, filters []*UploadFilter) (types.FileBrief, error) {
		return image.Thumb(file, width, height)
	}
}
