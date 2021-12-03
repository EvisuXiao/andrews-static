package transfer

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"time"
	"upload-test/config"
	"upload-test/pkg/utils"
)

const TYPE_COS = "cos"

type CosAdapter struct {
	cos *cos.Client
}

var cosAdapter = &CosAdapter{}

func NewCosAdapter() *CosAdapter {
	if !utils.IsEmpty(cosAdapter.cos) {
		return cosAdapter
	}
	u, _ := url.Parse(config.UploadConfig.Transfer.Addr)
	b := &cos.BaseURL{BucketURL: u}
	cosAdapter.cos = cos.NewClient(b, &http.Client{
		// 设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.UploadConfig.Transfer.Username,
			SecretKey: config.UploadConfig.Transfer.Password,
		},
	})
	return cosAdapter
}

func (a *CosAdapter) RemoteToLocal(filename string) error {
	return nil
}

func (a *CosAdapter) LocalToRemote(path, target string) error {
	return os.Rename(path, config.UploadFilePath(target))
}

func (a *CosAdapter) FileExists(filename string) bool {
	return false
}
