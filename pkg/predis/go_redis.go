package predis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	handle *redis.Client
	option *redis.Options
}

type Option func(*Service)

func New(opts ...Option) *Service {
	s := &Service{
		option: &redis.Options{},
	}

	for _, f := range opts {
		f(s)
	}

	return s
}

func (r *Service) Run() error {
	r.handle = redis.NewClient(r.option)

	err := r.handle.Get(context.Background(), "go-redis-testkey").Err()
	if errors.Is(err, redis.Nil) {
		return nil
	}

	return err
}

func (r *Service) GetConn() *redis.Client {
	return r.handle
}
