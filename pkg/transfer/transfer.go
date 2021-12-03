package transfer

import (
	"upload-test/config"
	"upload-test/pkg/logging"
)

type iTransfer interface {
	RemoteToLocal(string) error
	LocalToRemote(string, string) error
	FileExists(string) bool
}

var transferAdapter iTransfer

func Setup() {
	switch config.UploadConfig.Transfer.Type {
	case TYPE_LOCAL:
		transferAdapter = NewLocalAdapter()
	case TYPE_COS:
		transferAdapter = NewCosAdapter()
	default:
		logging.Fatal("Setup: transfer type not found", config.UploadConfig.Transfer.Type)
	}
}

func LocalToRemote(path, target string, overwrite bool) error {
	if !overwrite && transferAdapter.FileExists(path) {
		return nil
	}
	return transferAdapter.LocalToRemote(path, target)
}

func RemoteToLocal(filename string) error {
	return transferAdapter.RemoteToLocal(filename)
}
