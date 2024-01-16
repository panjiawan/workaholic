package plog

import (
	"go.uber.org/zap"
	"testing"
)

func TestDebug(t *testing.T) {
	Start("/tmp/", "debug", true, true)
	Debug("debug msg", zap.Int64("value", 12345))
}
