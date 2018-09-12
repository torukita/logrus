package helper

import (
	"github.com/sirupsen/logrus"
)

var DefaultLogLevel = logrus.InfoLevel

func LogLevel(level string) logrus.Level {
	v, err := logrus.ParseLevel(level)
	if err != nil {
		return DefaultLogLevel
	}
	return v
}
