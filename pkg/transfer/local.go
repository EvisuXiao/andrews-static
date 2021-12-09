package transfer

import (
	"os"

	"andrews-static/config"
)

const TYPE_LOCAL = "local"

type LocalAdapter struct{}

var localAdapter = &LocalAdapter{}

func NewLocalAdapter() *LocalAdapter {
	return localAdapter
}

func (a *LocalAdapter) RemoteToLocal(filename string) error {
	return nil
}

func (a *LocalAdapter) LocalToRemote(path, target string) error {
	return os.Rename(path, config.UploadFilePath(target))
}

func (a *LocalAdapter) FileExists(filename string) bool {
	return false
}
