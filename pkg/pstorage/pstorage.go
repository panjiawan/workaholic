package pstorage

import "io"

type IProvider interface {
	Init(*Config)
	PutFromStream(string, string, io.Reader, int64) (string, error)
}

type Service struct {
	p IProvider
}

type Config struct {
	SecretId      string
	SecretKey     string
	UseHTTPS      bool
	UseCdnDomains bool
}

type Option func(*Service)

type ProviderType int

var (
	Qiniu ProviderType = 1
)

func New(provider ProviderType, config *Config) *Service {
	s := &Service{}
	switch provider {
	case Qiniu:
		s.p = &QiniuCloud{}
		s.p.Init(config)
	}

	return s
}

func (s *Service) PutFromStream(bucket, path string, stream io.Reader, size int64) (string, error) {
	return s.p.PutFromStream(bucket, path, stream, size)
}
