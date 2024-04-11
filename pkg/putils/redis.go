package putils

import (
	"errors"
	"github.com/redis/go-redis/v9"
)

func RedisIsNil(err error) bool {
	return errors.Is(err, redis.Nil)
}
