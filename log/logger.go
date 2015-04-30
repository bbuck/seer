package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/seer-server/seer/errors"
)

type Logger interface {
	Log(LogLevel, string, ...interface{})
	Fatal(error, int)
}

type BasicLogger struct {
	*log.Logger
}

var (
	logMap       = make(map[string]Logger)
	logFlags     = log.Ldate | log.Ltime | log.Lshortfile
	specialFiles = map[string]io.Writer{
		":stdout:": os.Stdout,
		":stderr:": os.Stderr,
	}
)

func GetFileTarget(name string) io.Writer {
	if w, ok := specialFiles[name]; ok {
		return w
	} else {
		f, err := os.Open(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open %q as a source to log to: %s", name, err)

			os.Exit(errors.ErrFailedToCreateLogTarget)
		}

		return f
	}
}

func Make(name string, target io.Writer) Logger {
	if logger, ok := logMap[name]; ok {
		return logger
	} else {
		prefix := []string{"[", strings.Title(name), "] "}
		logger := &BasicLogger{log.New(target, strings.Join(prefix, ""), logFlags)}
		logMap[name] = logger

		return logger
	}
}

func Get(name string) (Logger, bool) {
	l, ok := logMap[name]

	return l, ok
}

func (l *BasicLogger) Fatal(err error, exitCode int) {
	l.Log(LogError, err.Error())
	os.Exit(exitCode)
}

func (l *BasicLogger) Log(level LogLevel, format string, args ...interface{}) {
	if level <= MaxLogLevel {
		l.Logger.Printf("%10s   %s", level, fmt.Sprintf(format, args...))
	}
}
