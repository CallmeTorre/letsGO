package option_a

import (
	"fmt"
	"os"
	"strings"

	"github.com/CallmeTorre/letsGO/api/config"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}
	Log = &logrus.Logger{
		Level:     level,
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
	}
}

func Info(message string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Info(message)
}

func Debug(message string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Debug(message)
}

func Error(message string, err error, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	message = fmt.Sprintf("%s - ERROR - %s", message, err.Error())
	Log.WithFields(parseFields(tags...)).Error(message)
}

func Fatal(message string, err error, tags ...string) {
	if Log.Level < logrus.FatalLevel {
		return
	}
	message = fmt.Sprintf("%s - ERROR - %s", message, err.Error())
	Log.WithFields(parseFields(tags...)).Fatal(message)
}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		elem := strings.Split(tag, ":")
		result[strings.TrimSpace(elem[0])] = strings.TrimSpace(elem[1])
	}
	return result
}
