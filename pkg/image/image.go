package image

import (
	"fmt"
	"github.com/disintegration/imaging"
	"upload-test/pkg/utils"
	"upload-test/types"
)

var DefaultFilter = imaging.Lanczos

func Resize(file types.FileBrief, width, height int) (types.FileBrief, error) {
	img, err := imaging.Open(file.FilePath)
	if !utils.IsEmpty(err) {
		return types.FileBrief{}, err
	}
	newImg := imaging.Resize(img, width, height, DefaultFilter)
	file.Basename = fmt.Sprintf("%s_%dx%d", file.Basename, width, height)
	file.FormatFromBasename()
	return file, imaging.Save(newImg, file.FilePath)
}

func Thumb(file types.FileBrief, width, height int) (types.FileBrief, error) {
	img, err := imaging.Open(file.FilePath)
	if !utils.IsEmpty(err) {
		return types.FileBrief{}, err
	}
	thumb := imaging.Thumbnail(img, width, height, DefaultFilter)
	file.Basename = fmt.Sprintf("%s_thumb_%dx%d", file.Basename, width, height)
	file.FormatFromBasename()
	return file, imaging.Save(thumb, file.FilePath)
}
