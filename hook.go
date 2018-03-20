package delog

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var (
	_ logrus.Hook = new(Hook)
)

type Hook struct {
	writer    io.Writer
	formatter Formatter
	logLevels []logrus.Level
}

func NewHook(formatter logrus.Formatter, wr io.Writer, logLevels ...logrus.Level) *Hook {
	if len(logLevels) == 0 {
		logLevels = []logrus.Level{logrus.DebugLevel}
	}
	if formatter == nil {
		formatter = &logrus.TextFormatter{FullTimestamp: true}
	}
	if wr == nil {
		wr = os.Stdout
	}
	delogFormatter := *NewFormatter(formatter)
	delogFormatter.stackOffset = 2
	return &Hook{
		writer:    wr,
		formatter: delogFormatter,
		logLevels: logLevels,
	}
}

func (hook *Hook) Levels() []logrus.Level {
	return hook.logLevels
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	msg, err := hook.formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.writer.Write(msg)
	return err
}
