package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	dirPrefix         = "logs"
	timestampFieldKey = "@timestamp"
	messageFieldKey   = "message"

	logPathFormat     = "%s/%s"
	logFileNameFormat = "%s/%s.txt"

	errorCreatingFile = "failed to initialize log file"
)

// Logger interface abstracts the logger package and offers an interface with three kind of logger
type Logger interface {
	Info(ctx context.Context, msg string, fields log.Fields)
	Warn(ctx context.Context, msg string, fields log.Fields)
	Error(ctx context.Context, err error, msg string, fields log.Fields)
	Debug(ctx context.Context, msg string, fields log.Fields)
}

// NewLogger returns an implementation of Logger
func NewLogger(debugMode bool) Logger {
	outputFile := newLogFile(time.Now().UTC(), getLogPath())

	logrus := log.New()
	logrus.SetOutput(outputFile)

	logrus.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime: timestampFieldKey,
			log.FieldKeyMsg:  messageFieldKey,
		},
	})

	setLogLevel(logrus, debugMode)

	return &logger{
		logrus: logrus,
	}
}

func setLogLevel(logrus *log.Logger, debugMode bool) {
	logLevel := log.TraceLevel
	if debugMode {
		logLevel = log.DebugLevel
	}
	logrus.SetLevel(logLevel)
}

// logger implements Logger interface
type logger struct {
	logrus *log.Logger
}

func (l logger) Info(ctx context.Context, msg string, fields log.Fields) {
	l.logrus.WithContext(ctx).WithFields(fields).Info(msg)
}

func (l logger) Warn(ctx context.Context, msg string, fields log.Fields) {
	l.logrus.WithContext(ctx).WithFields(fields).Warn(msg)
}

func (l logger) Error(ctx context.Context, err error, msg string, fields log.Fields) {
	l.logrus.WithContext(ctx).WithError(err).WithFields(fields).Error(msg)
}

func (l logger) Debug(ctx context.Context, msg string, fields log.Fields) {
	l.logrus.WithContext(ctx).WithFields(fields).Debug(msg)
}

// newLogFile creates file with the current date
func newLogFile(t time.Time, logPath string) io.Writer {
	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		log.WithError(err).Warn(errorCreatingFile)
		return os.Stdout
	}

	fileName := fmt.Sprintf(logFileNameFormat, logPath, t.UTC().Format("01-02-2006"))
	logFile, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	return logFile
}

func getLogPath() string {
	return fmt.Sprintf(logPathFormat, getProjectRootDirectory(), dirPrefix)
}

func getProjectRootDirectory() string {
	currDir, _ := os.Getwd()
	return filepath.Dir(filepath.Dir(currDir))
}
