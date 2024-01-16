package plog

import "go.uber.org/zap/zapcore"

func Debug(msg string, fields ...zapcore.Field) {
	zapLog.Debug(msg, fields...)
}
