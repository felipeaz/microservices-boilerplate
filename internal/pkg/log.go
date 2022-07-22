package pkg

import (
	"log"
	"os"
)

const (
	infoPrefix  = "[INFO]"
	warnPrefix  = "[WARNING]"
	errPrefix   = "[ERROR]"
	debugPrefix = "[DEBUG]"
	logPath     = "logs.txt"
)

// Logger interface abstracts the log package and offers an interface with three kind of log
type Logger interface {
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Debug(v ...interface{})
}

// NewLogger returns an implementation of Logger
func NewLogger(debugMode bool) Logger {
	lf, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("unable to initialize log file", logPath, err)
	}
	l := logger{
		info: log.New(os.Stdout, infoPrefix, log.LstdFlags|log.Lshortfile),
		warn: log.New(os.Stdout, warnPrefix, log.LstdFlags|log.Lshortfile),
		err:  log.New(lf, errPrefix, log.LstdFlags|log.Llongfile),
	}
	if debugMode {
		l.debug = log.New(os.Stdout, debugPrefix, log.LstdFlags|log.Lshortfile)
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
