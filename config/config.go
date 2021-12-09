package config

import (
	"github.com/EvisuXiao/andrews-common/config"
)

const ServiceName = "restapi-static-andrews"

func Init() {
	config.Init(ServiceName)
}