package main

import (
	"fmt"
	"net/http"
	"os"
	"upload-test/config"
	"upload-test/pkg/logging"
	"upload-test/pkg/transfer"
	"upload-test/pkg/translation"
	"upload-test/pkg/utils"
	"upload-test/router"
)

func init() {
	config.Setup()
}

func main() {
	transfer.Setup()
	translation.Setup()
	logging.Info("All Environment setup successfully!")
	logging.Info("The process id is %d", os.Getpid())
	startHttpServer()
}

func startHttpServer() {
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.ServerConfig.HttpPort),
		Handler:        router.InitRouter(),
		ReadTimeout:    config.ServerConfig.Timeout.Read,
		WriteTimeout:   config.ServerConfig.Timeout.Write,
		MaxHeaderBytes: 1 << 20,
	}
	logging.Info("Start HTTP server with listening port %d", config.ServerConfig.HttpPort)
	logging.Info("Service %s is running!", config.ServiceName)
	if err := server.ListenAndServe(); !utils.IsEmpty(err) {
		logging.Fatal("Start HTTP server failed with err: %+v", err)
	}
}
