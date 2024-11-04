package predis

import (
	"fmt"
	"time"
)

func WithConnection(host string, port int) Option {
	return func(s *Service) {
		s.option.Addr = fmt.Sprintf("%s:%d", host, port)
	}
}

func WithAuth(auth string) Option {
	return func(s *Service) {
		s.option.Password = auth
	}
}

func WithLimit(minIdle, maxIdle int) Option {
	return func(s *Service) {
		s.option.MinIdleConns = minIdle
		s.option.MaxIdleConns = maxIdle
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(s *Service) {
		s.option.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *Service) {
		s.option.WriteTimeout = timeout
	}
}

func WithDialTimeout(timeout time.Duration) Option {
	return func(s *Service) {
		s.option.DialTimeout = timeout
	}
}

func WithDB(db int) Option {
	return func(s *Service) {
		s.option.DB = db
	}
}
