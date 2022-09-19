package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"microservices-boilerplate/internal/pkg/dir"
)

const (
	infoPrefix  = "[INFO]"
	warnPrefix  = "[WARNING]"
	errPrefix   = "[ERROR]"
	debugPrefix = "[DEBUG]"
	dirPrefix   = "logs"
)

// Logger interface abstracts the log package and offers an interface with three kind of log
type Logger interface {
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Debug(v ...interface{})
}

// NewLogger returns an implementation of Logger
func NewLogger(output io.Writer, debugMode bool) Logger {
	l := logger{
		info: log.New(output, infoPrefix, log.LstdFlags|log.Lshortfile),
		warn: log.New(output, warnPrefix, log.LstdFlags|log.Lshortfile),
		err:  log.New(output, errPrefix, log.LstdFlags|log.Lshortfile),
	}
	if debugMode {
		l.debug = log.New(output, debugPrefix, log.LstdFlags|log.Lshortfile)
	}
	return l
}

// logger implements Logger interface
type logger struct {
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
	debug *log.Logger
}

func (l logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

func (l logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l logger) Debug(v ...interface{}) {
	if l.debug != nil {
		l.debug.Println(v...)
	}
}

// NewLogFile creates file with the current date
func NewLogFile(t time.Time, logPath string) io.Writer {
	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		log.Println("failed to initialize log file", err)
		return os.Stdout
	}

	fileName := fmt.Sprintf("%s/%s.txt", logPath, t.UTC().Format("01-02-2006"))
	logFile, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	return logFile
}

func GetLogPath() string {
	return fmt.Sprintf("%s/%s", dir.GetProjectRootDirectory(), dirPrefix)
}
