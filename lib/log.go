package lib

import "go.uber.org/zap"

var log *zap.Logger

func InitLog() {
	log, _ = zap.NewProduction()
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
