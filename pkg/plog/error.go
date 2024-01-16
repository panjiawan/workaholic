package plog

import "go.uber.org/zap/zapcore"

func Error(msg string, fields ...zapcore.Field) {
	zapLog.Error(msg, fields...)
}
