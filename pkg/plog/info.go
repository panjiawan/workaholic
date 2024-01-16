package plog

import "go.uber.org/zap/zapcore"

func Info(msg string, fields ...zapcore.Field) {
	zapLog.Info(msg, fields...)
}
