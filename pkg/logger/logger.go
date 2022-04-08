package logger

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Logrus      *logrus.Logger
	ContextKeys []string
}

//NewLogger will return a logger instance. It accepts contextKeys array of keys which need
// to fetched from context if user want to log with context
func NewLogger(contextKeys []string) *Logger {
	return &Logger{
		Logrus:      logrus.New(),
		ContextKeys: contextKeys,
	}
}

//WithContext will accept the context and will return the instance of ContextLogger
func (l Logger) WithContext(ctx context.Context) *logrus.Entry {

	cl := l.Logrus.WithTime(time.Now())

	for _, contextKey := range l.ContextKeys {
		contextValue, ok := ctx.Value(contextKey).(string)
		if !ok {
			contextValue = "not tracked"
		}
		cl = cl.WithField(contextKey, contextValue)
	}

	return cl
}
