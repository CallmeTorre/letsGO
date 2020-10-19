package option_b

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	Log, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func Info(message string, tags ...zap.Field) {
	Log.Info(message, tags...)
	Log.Sync()
}

func Debug(message string, tags ...zap.Field) {
	Log.Debug(message, tags...)
	Log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	message = fmt.Sprintf("%s - ERROR - %s", message, err.Error())
	Log.Error(message, tags...)
	Log.Sync()
}

func Fatal(message string, err error, tags ...zap.Field) {
	message = fmt.Sprintf("%s - ERROR - %s", message, err.Error())
	Log.Fatal(message, tags...)
	Log.Sync()
}
