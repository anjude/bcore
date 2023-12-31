package logger

import (
	"github.com/sirupsen/logrus"
)

type BeanLogger struct {
	*logrus.Logger
}

func Default() *BeanLogger {
	return &BeanLogger{
		Logger: logrus.New(),
	}
}
