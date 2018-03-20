# delog

Magic logrus debug formatter and hook. Adds prefixes a-la [FILENAME:FUNCTION:LINE] to log messages.

Formatter usage:

```go
    package main

    import (
        "github.com/ninedraft/delog"
        "github.com/sirupsen/logrus"
    )

    func main() {
        log := logrus.New()
        log.SetLevel(logrus.DebugLevel)
        log.Formatter = delog.NewFormatter(&logrus.TextFormatter{})
        log.Debugf("debug message") 
        // --> time="2018-03-20T19:30:23+03:00" level=debug msg="[path/to/your/package/main.go:main:12] debug message"
    }
```

Hook usage:

```go
    package main

    import (
        "os"

        "github.com/ninedraft/delog"
        "github.com/sirupsen/logrus"
    )

    func main() {
        log := logrus.New()
        log.SetLevel(logrus.DebugLevel)
        log.AddHook(delog.NewHook(&logrus.TextFormatter{}, os.Stdout /* or file */))
        log.Debugf("debug message") 
        // --> time="2018-03-20T19:30:23+03:00" level=debug msg="[path/to/your/package/main.go:main:12] debug message"
    }

```