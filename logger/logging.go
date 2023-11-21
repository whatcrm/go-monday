package logging

import (
	"fmt"
	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"time"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	newHook, is := refreshLogFile()
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

func (l *Logger) Field(v interface{}) Logger {
	return Logger{l.WithField("key", v)}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)

	allFile, err := CreateFile()
	if err != nil {
		fmt.Println(err)
	}

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
			filename := path.Base(frame.File)
			return fmt.Sprintf(" %s:[%d]", filename, frame.Line)
		},
	}

	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})
	l.SetLevel(logrus.TraceLevel)
	e = logrus.NewEntry(l)
}

func CreateFile() (*os.File, error) {
	today := time.Now()

	logFilename = today.Format("2006-01-02")

	filename := "./logs/graphql/" + today.Format("2006-01-02") + ".log"
	allFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	return allFile, err
}

func refreshLogFile() (*writerHook, bool) {
	today := time.Now()
	fileDate := today.Format("2006-01-02")

	if fileDate != logFilename {
		logFilename = fileDate

		fileDate := "logs/" + fileDate + ".log"
		allFile, err := os.OpenFile(fileDate, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
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
