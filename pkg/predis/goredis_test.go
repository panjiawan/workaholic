package predis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestConn(t *testing.T) {
	service := New(WithConnection("127.0.0.1", 6379))
	err := service.Run()
	t.Log("res", err, "--")
	slice := service.GetConn().ZRangeByScore(context.Background(), "test_zsort", &redis.ZRangeBy{
		Min:    fmt.Sprintf("%d", 2),
		Max:    fmt.Sprintf("%d", 5),
		Offset: int64(0),
		Count:  int64(20),
	})
	t.Log(slice.Result())
}
