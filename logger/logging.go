package logging

import (
	"fmt"
	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"time"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	newHook, is := refreshLogFile(entry)
	if is {
		*hook = *newHook
	}

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
var logFilename string

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
	l.SetLevel(logrus.TraceLevel)
	e = logrus.NewEntry(l)
}

func refreshLogFile(entry *logrus.Entry) (*writerHook, bool) {
	today := time.Now()
	fileDate := "/gql-" + today.Format("2006-01-02") + ".log"

	if fileDate != logFilename {
		folderName, ok := entry.Data["domain"].(string)
		if !ok {
			folderName = "core"
		}

		folder := "./logs/" + folderName
		err := ensureFolder(folder)
		if err != nil {
			return nil, false
		}

		logFilename = folder + fileDate
		allFile, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic(err)
		}

		return &writerHook{
			Writer:    []io.Writer{allFile, os.Stdout},
			LogLevels: logrus.AllLevels,
		}, true
	}

	return nil, false
}

func ensureFolder(path string) error {
	_, err := os.Stat(path)

	switch {
	case os.IsNotExist(err):
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("error creating folder: %v", err)
		}
	case err != nil:
		return fmt.Errorf("error checking folder existence: %v", err)
	}

	return nil
}
