package main

import (
	"os"

	"github.com/EvisuXiao/andrews-common/logging"
	cHttp "github.com/EvisuXiao/andrews-common/server/http"

	"andrews-static/config"
	"andrews-static/pkg/transfer"
	"andrews-static/router"
)

func init() {
	config.Init()
}

func main() {
	transfer.Init()
	logging.Info("All Environment setup successfully!")
	logging.Info("The process id is %d", os.Getpid())
	cHttp.StartServer(router.Init())
}
