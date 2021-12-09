package config

import (
	"strconv"
	"strings"

	"github.com/EvisuXiao/andrews-common/config"
)

type Upload struct {
	ImageMaxSizeStr string   `json:"image_max_size" default:"5m"`
	VideoMaxSizeStr string   `json:"video_max_size" default:"50m"`
	AudioMaxSizeStr string   `json:"audio_max_size" default:"10m"`
	ImageMaxSize    int64    `json:"-"`
	VideoMaxSize    int64    `json:"-"`
	AudioMaxSize    int64    `json:"-"`
	Transfer        Transfer `json:"transfer"`
}
type Transfer struct {
	Type     string `json:"type" default:"local"`
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
	Path     string `json:"path" default:"/data/static/"`
}

var UploadConfig = &Upload{}

func init() {
	config.RegisterConfig(UploadConfig)
}

func (c *Upload) Name() string {
	return "upload"
}

func (c *Upload) Source() string {
	return ""
}

func (c *Upload) Check() error {
	return nil
}

func (c *Upload) Init() {
	c.ImageMaxSize = sizeFormatter(c.ImageMaxSizeStr)
	c.VideoMaxSize = sizeFormatter(c.VideoMaxSizeStr)
	c.AudioMaxSize = sizeFormatter(c.AudioMaxSizeStr)
}

func sizeFormatter(size string) int64 {
	if size == "" {
		return 0
	}
	numStr := size[:len(size)-1]
	num, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		return 0
	}
	unit := strings.ToLower(size[len(size)-1:])
	switch unit {
	case "b":
		return num
	case "k":
		return num * 1024
	case "m":
		return num * 1024 * 1024
	case "g":
		return num * 1024 * 1024 * 1024
	default:
		return 0
	}
}

func UploadFilePath(filename string) string {
	return UploadConfig.Transfer.Path + filename
}
