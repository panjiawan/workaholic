package pstorage

import (
	"context"
	"io"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiniuCloud struct {
	accessKey     string
	secretKey     string
	useHTTPS      bool
	useCdnDomains bool
}

func (q *QiniuCloud) Init(config *Config) {
	q.accessKey = config.SecretId
	q.secretKey = config.SecretKey
	q.useHTTPS = config.UseHTTPS
	q.useCdnDomains = config.UseCdnDomains
}

func (q *QiniuCloud) PutFromStream(bucket, path string, stream io.Reader, size int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(q.accessKey, q.secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 是否使用https域名
	cfg.UseHTTPS = q.useHTTPS
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = q.useCdnDomains
	formUploader := storage.NewFormUploader(&cfg)
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	ret := storage.PutRet{}
	err := formUploader.Put(context.Background(), &ret, upToken, path, stream, size, &putExtra)
	if err != nil {
		return "", err
	}
	return ret.Hash, nil
}
