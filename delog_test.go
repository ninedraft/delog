package delog

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"testing"
)

func run() string {
	log := logrus.New()
	buf := &bytes.Buffer{}
	log.SetLevel(logrus.DebugLevel)
	log.Formatter = &Formatter{
		formatter: &logrus.TextFormatter{
			FullTimestamp: true,
		},
	}
	log.Out = buf
	log.Debugf("test message")
	return buf.String()
}
func TestFormatter(test *testing.T) {
	test.Logf("%q", run())
}

func TestHook(test *testing.T) {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.AddHook(NewHook(nil, nil))
	log.Debugf("test message")
}
