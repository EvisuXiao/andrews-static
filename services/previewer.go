package services

import (
	"fmt"
	"upload-test/config"
	"upload-test/types"
)

type Previewer struct {
	file types.FileBrief
}

func NewPreviewer(filename string, mediaType types.MEDIA_TYPE) *Previewer {
	previewer := &Previewer{}
	previewer.file.FilePath = config.UploadFilePath(fmt.Sprintf("%s/%s", mediaType, filename))
	previewer.file.FormatFromFilepath()
	return previewer
}

func (p *Previewer) Preview() string {
	return p.file.FilePath
}
