package types

import (
	"fmt"
	"path/filepath"
	"strings"
)

type MEDIA_TYPE string

const (
	MEDIA_TYPE_IMAGE   = MEDIA_TYPE("image")
	MEDIA_TYPE_GIF     = MEDIA_TYPE("gif")
	MEDIA_TYPE_VIDEO   = MEDIA_TYPE("video")
	MEDIA_TYPE_AUDIO   = MEDIA_TYPE("audio")
	MEDIA_TYPE_FILE    = MEDIA_TYPE("file")
	MEDIA_TYPE_UNKNOWN = MEDIA_TYPE("unknown")
)

type FileBrief struct {
	Filename  string
	FilePath  string
	Basename  string
	BaseDir   string
	Ext       string
	MediaType MEDIA_TYPE
}

func (f *FileBrief) FormatFromBasename() {
	f.Filename = fmt.Sprintf("%s.%s", f.Basename, f.Ext)
	f.FilePath = f.BaseDir + f.Filename
	f.FormatMediaType()
}

func (f *FileBrief) FormatFromFilename() {
	fArr := strings.Split(f.Filename, ".")
	fLen := len(fArr)
	if fLen > 1 {
		f.Basename = strings.Join(fArr[:fLen-1], ".")
		f.Ext = strings.ToLower(fArr[fLen-1])
	} else {
		f.Basename = f.Filename
	}
	f.FilePath = f.BaseDir + f.Filename
	f.FormatMediaType()
}

func (f *FileBrief) FormatFromFilepath() {
	f.BaseDir = filepath.Dir(f.FilePath) + "/"
	f.Filename = filepath.Base(f.FilePath)
	f.FormatFromFilename()
}

func (f *FileBrief) FormatMediaType() {
	switch f.Ext {
	case "":
		f.MediaType = MEDIA_TYPE_UNKNOWN
	case "jpg", "jpeg", "png":
		f.MediaType = MEDIA_TYPE_IMAGE
	case "gif":
		f.MediaType = MEDIA_TYPE_GIF
	case "mp4", "mpeg":
		f.MediaType = MEDIA_TYPE_VIDEO
	case "mp3", "wav", "m4a":
		f.MediaType = MEDIA_TYPE_AUDIO
	default:
		f.MediaType = MEDIA_TYPE_FILE
	}
}

func (f FileBrief) InvalidMedia() bool {
	return f.MediaType == MEDIA_TYPE_UNKNOWN
}
