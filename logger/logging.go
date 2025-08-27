package logging

import (
	"io"
	"os"
	"runtime"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) Field(v string) Logger {
	return Logger{l.WithField("domain", v)}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)

	l.SetOutput(io.Discard)
	l.Formatter = &formatter.Formatter{
		FieldsOrder:      nil,
		TimestampFormat:  "2006 02 Jan | 15:04:05 UTC",
		HideKeys:         true,
		NoColors:         true,
		NoFieldsColors:   false,
		NoFieldsSpace:    false,
		ShowFullLevel:    true,
		NoUppercaseLevel: false,
		TrimMessages:     false,
		CallerFirst:      false,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			return ""
		},
	}

	l.AddHook(&writerHook{
		LogLevels: logrus.AllLevels,
	})

	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.TraceLevel)
	e = logrus.NewEntry(l)
}
