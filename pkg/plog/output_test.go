package plog

import (
	"go.uber.org/zap"
	"testing"
)

func TestInfo(t *testing.T) {
	Start("/tmp/", "test.log", true, false)
	Info("test msg", zap.Any("1", "2"))
}
