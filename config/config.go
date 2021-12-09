package config

import (
	"strconv"
	"strings"

	"github.com/EvisuXiao/andrews-common/config"
)

const ServiceName = "restapi-static-andrews"

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

func Init() {
	config.Init(ServiceName)
	loadConf()
	initConf()
}

func loadConf() {
	config.MapTo("upload.json", UploadConfig)
}

func initConf() {
	UploadConfig.ImageMaxSize = sizeFormatter(UploadConfig.ImageMaxSizeStr)
	UploadConfig.VideoMaxSize = sizeFormatter(UploadConfig.VideoMaxSizeStr)
	UploadConfig.AudioMaxSize = sizeFormatter(UploadConfig.AudioMaxSizeStr)
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
