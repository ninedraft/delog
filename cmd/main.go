package main

import (
	"github.com/ninedraft/delog"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.Formatter = delog.NewFormatter(&logrus.TextFormatter{})
	log.Debugf("test message")
}
