package config

import (
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// App default configurations
var (
	DefaultLoglvl    = logrus.InfoLevel
	DefaultLogOutput = os.Stdout
	DefaultFormater  = &logrus.TextFormatter{}
)

func ParseLogLevel(level string) logrus.Level {
	switch level {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}

// parse Formatter from string
func ParseLogOutput(lvl string) (io.Writer, error) {
	switch lvl {
	case "terminal":
		return os.Stdout, nil
	case "text-file":
		file, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return nil, errors.Wrap(err, " while creating/opening text file")
		}
		return file, nil
	default:
		return DefaultLogOutput, nil
	}
}
