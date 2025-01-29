package logger

import "go.uber.org/zap"

var Log *zap.Logger

func InitLogger() {
	if Log != nil {
		return
	}

	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	Log = logger
}

func GetLogger() *zap.Logger {
	if Log == nil {
		InitLogger()
	}
	return Log
}
