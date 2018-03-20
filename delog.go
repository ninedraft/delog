package delog

import (
	"github.com/sirupsen/logrus"
)

var (
	_ logrus.Formatter = new(Formatter)
)

type Formatter struct {
	formatter logrus.Formatter
}

func (dbgFormatter *Formatter) Format(entry *logrus.Entry) ([]byte, error) {

	msg, err := dbgFormatter.Format(entry)
	return msg, err
}
