package logger

import (
	"github.com/sirupsen/logrus"
)

type BeanLogger struct {
	*logrus.Logger
}
