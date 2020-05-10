package common

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

type Logger struct {
	*logging.Logger
}

var Log = &Logger{Logger: logging.MustGetLogger("om")}

func InitLog() {
	logBackend := logging.NewBackendFormatter(
		logging.NewLogBackend(os.Stdout, "", 0),
		logging.MustStringFormatter(
			`%{time:2006-01-02 15:04:05.000} %{level:.4s} %{shortfile} %{message:.1024s}`,
		))

	levelBackend := logging.SetBackend(logBackend)
	levelBackend.SetLevel(logging.DEBUG, "")
	Log.SetBackend(levelBackend)
}

func (l *Logger) Log(v ...interface{}) {
	l.Info(v...)
}

func (l *Logger) Logf(format string, v ...interface{}) {
	l.Infof(format, v...)
}

// 给sql用
func (l *Logger) Print(v ...interface{}) {
	msg := gorm.LogFormatter(v...)
	l.Debug(fmt.Sprintf("%v,%v,%v", msg[0], msg[1], msg[2]))
}
