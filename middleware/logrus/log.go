package logrus

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {

	logrus.SetOutput(os.Stdout)
	//设置日志等级
	logrus.SetLevel(logrus.DebugLevel)
}

func Info(template string, args ...interface{}) {
	logrus.Infof(template, args)
}

func Debug(template string, args ...interface{}) {
	logrus.Debugf(template, args)
}

func Warn(template string, args ...interface{}) {
	logrus.Warnf(template, args)
}

func Error(template string, args ...interface{}) {
	logrus.Errorf(template, args)
}
