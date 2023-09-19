package logger

import (
	"github.com/sirupsen/logrus"
	"tinkapi/v2/pkg/config"
)

func NewLogger(env string) *logrus.Logger { // TODO: different logger levels on .ENV
	logger := &logrus.Logger{}
	if env == config.ENV_LOCAL {
		logger = &logrus.Logger{
			// Formatter:  &logrus.JSONFormatter{},
			Level: logrus.DebugLevel,
		}
	} else if env == config.ENV_DEV {
		logger = &logrus.Logger{
			Formatter: &logrus.JSONFormatter{},
			Level:     logrus.DebugLevel,
		}
	} else if env == config.ENV_PROD {
		logger = &logrus.Logger{
			Formatter: &logrus.JSONFormatter{},
			Level:     logrus.InfoLevel,
		}
	} else {
		logger.Fatalf("invalid config env parameter: got %s, expected %s, %s or %s",
			env, config.ENV_LOCAL, config.ENV_DEV, config.ENV_PROD)
	}

	return logger
}
