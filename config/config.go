package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ServiceName = "restapi.uploadtest"

	EnvLocal   = "local"
	EnvTesting = "testing"
	EnvProd    = "production"

	tempPath = "runtime/"
)

var dir string

type Server struct {
	Env      string  `json:"env"`
	HttpPort int     `json:"http_port"`
	Timeout  Timeout `json:"timeout"`
}
type Timeout struct {
	Read  time.Duration `json:"read"`
	Write time.Duration `json:"write"`
}

var ServerConfig = &Server{}

type Upload struct {
	ImageMaxSizeStr string   `json:"image_max_size"`
	VideoMaxSizeStr string   `json:"video_max_size"`
	AudioMaxSizeStr string   `json:"audio_max_size"`
	ImageMaxSize    int64    `json:"-"`
	VideoMaxSize    int64    `json:"-"`
	AudioMaxSize    int64    `json:"-"`
	Transfer        Transfer `json:"transfer"`
}
type Transfer struct {
	Type     string `json:"type"`
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
	Path     string `json:"path"`
}

var UploadConfig = &Upload{}

func Setup() {
	parseFlag()
	loadConf()
	initConf()
}

func parseFlag() {
	flag.StringVar(&dir, "dir", "./", "The application directory")
	flag.Parse()
	dir = addDirSlash(dir)
}

func loadConf() {
	log.Println("[INFO] Load configuration")
	mapTo("server.json", ServerConfig)
	mapTo("upload.json", UploadConfig)
}

func initConf() {
	if ServerConfig.Env != EnvLocal && ServerConfig.Env != EnvProd {
		ServerConfig.Env = EnvTesting
	}
	ServerConfig.Timeout.Read = ServerConfig.Timeout.Read * time.Second
	ServerConfig.Timeout.Write = ServerConfig.Timeout.Write * time.Second
	UploadConfig.ImageMaxSize = sizeFormatter(UploadConfig.ImageMaxSizeStr)
	UploadConfig.VideoMaxSize = sizeFormatter(UploadConfig.VideoMaxSizeStr)
	UploadConfig.AudioMaxSize = sizeFormatter(UploadConfig.AudioMaxSizeStr)
	UploadConfig.Transfer.Path = addDirSlash(UploadConfig.Transfer.Path)
}

func mapTo(section string, config interface{}) {
	filename := AppFilePath(fmt.Sprintf("conf/%s", section))
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatalf("[FATAL] Setup fatal: open %s error: %v\n", filename, err)
	}
	read, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("[FATAL] Setup fatal: read %s error: %v\n", filename, err)
	}
	err = json.Unmarshal(read, config)
	if err != nil {
		log.Fatalf("[FATAL] Setup fatal: parse %s error: %v\n", filename, err)
	}
	log.Printf("[INFO] %s config setup successfully!", section)
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

func addDirSlash(path string) string {
	if path == "" {
		path = "."
	}
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	return path
}

func AppFilePath(filename string) string {
	return dir + filename
}

func UploadFilePath(filename string) string {
	return UploadConfig.Transfer.Path + filename
}

func TempFilePath(filename string) string {
	return AppFilePath(tempPath + filename)
}

func IsLocalEnv() bool {
	return ServerConfig.Env == EnvLocal
}

func IsTestingEnv() bool {
	return ServerConfig.Env == EnvTesting
}

func IsProdEnv() bool {
	return ServerConfig.Env == EnvProd
}
