package pkg

import "github.com/stretchr/testify/mock"

type Logger struct {
	mock.Mock
}

func (l *Logger) Info(v ...interface{}) {
	l.Mock.Called(v)
}

func (l *Logger) Warn(v ...interface{}) {
	l.Mock.Called(v)
}

func (l *Logger) Error(v ...interface{}) {
	l.Mock.Called(v)
}

func (l *Logger) Debug(v ...interface{}) {
	l.Mock.Called(v)
}
