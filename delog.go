package delog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	gopaths = []string{}
)

var (
	_ logrus.Formatter = new(Formatter)
)

type Formatter struct {
	formatter logrus.Formatter
}

func NewFormatter(formatter logrus.Formatter) *Formatter {
	if formatter == nil {
		formatter = &logrus.TextFormatter{}
	}
	return &Formatter{
		formatter: formatter,
	}
}

func (dbgFormatter *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	skip := 0
	if len(entry.Data) == 0 {
		skip = 6
	} else {
		skip = 4
	}
	caller, filePath, line, _ := runtime.Caller(skip)
	frame, _ := runtime.CallersFrames([]uintptr{caller}).Next()
	file := stripGopath(filePath)
	fnName := strings.Split(path.Base(frame.Function), ".")[1]

	ddbgData := fmt.Sprintf("%s:%s:%d", file, fnName, line)
	entry.Message = fmt.Sprintf("[%s] %s", ddbgData, entry.Message)
	msg, err := dbgFormatter.formatter.Format(entry)
	return msg, err
}

func stripGopath(str string) string {
	str = sanitizePath(str)
	for _, gopath := range gopaths {
		if strings.HasPrefix(str, gopath) {
			return str[len(gopath):]
		}
	}
	return str
}

func sanitizePath(str string) string {
	str = filepath.ToSlash(str)
	probablyDiskAndGopath := strings.SplitN(str, ":", 2)
	if len(probablyDiskAndGopath) > 1 {
		str = strings.ToLower(probablyDiskAndGopath[0]) + ":" + probablyDiskAndGopath[1]
	}
	return str
}

func init() {
	goroot := runtime.GOROOT()
	gopaths = []string{path.Join(filepath.ToSlash(goroot), "src") + "/"}
	for _, gopath := range filepath.SplitList(os.Getenv("GOPATH")) {
		if gopath != "" {
			gopath = sanitizePath(gopath)
			gopaths = append(gopaths, path.Join(gopath, "src")+"/")
		}
	}
}
