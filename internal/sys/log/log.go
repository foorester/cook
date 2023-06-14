package log

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

type LogLevel int

const (
	Error LogLevel = iota
	Info
	Debug
)

type Logger interface {
	SetLogLevel(level LogLevel)
	Debug(v ...any)
	Debugf(format string, a ...any)
	Info(v ...any)
	Infof(format string, a ...any)
	Error(v ...any)
	Errorf(format string, a ...any)
}

type SimpleLogger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger
	logLevel    LogLevel
}

func NewLogger(logLevel string) *SimpleLogger {
	level := toValidLevel(logLevel)
	return &SimpleLogger{
		debugLogger: log.New(os.Stdout, "[DBG] ", log.LstdFlags),
		infoLogger:  log.New(os.Stdout, "[INF] ", log.LstdFlags),
		errorLogger: log.New(os.Stderr, "[ERR] ", log.LstdFlags),
		logLevel:    level,
	}
}

func (l *SimpleLogger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

func (l *SimpleLogger) Debug(v ...any) {
	if l.logLevel <= Debug {
		l.debugLogger.Println(v...)
	}
}

func (l *SimpleLogger) Debugf(format string, a ...any) {
	if l.logLevel <= Debug {
		message := fmt.Sprintf(format, a...)
		l.debugLogger.Println(message)
	}
}

func (l *SimpleLogger) Info(v ...any) {
	if l.logLevel <= Info {
		l.infoLogger.Println(v...)
	}
}

func (l *SimpleLogger) Infof(format string, a ...any) {
	if l.logLevel <= Info {
		message := fmt.Sprintf(format, a...)
		l.infoLogger.Println(message)
	}
}

func (l *SimpleLogger) Error(v ...any) {
	if l.logLevel <= Error {
		l.errorLogger.Println(v...)
	}
}

func (l *SimpleLogger) Errorf(format string, a ...any) {
	if l.logLevel <= Error {
		message := fmt.Sprintf(format, a...)
		l.errorLogger.Println(message)
	}
}

func toValidLevel(level string) LogLevel {
	level = strings.ToLower(level)

	switch level {
	case "debug", "dbg":
		return Debug
	case "info", "inf":
		return Info
	case "error", "err":
		return Error
	default:
		return Error
	}
}

// SetDebugOutput set the internal logger.
// Used for package testing.
func (sl *SimpleLogger) SetDebugOutput(debug *bytes.Buffer) {
	sl.debugLogger = log.New(debug, "", 0)
}

// SetInfoOutput set the internal logger.
// Used for package testing.
func (sl *SimpleLogger) SetInfoOutput(info *bytes.Buffer) {
	sl.infoLogger = log.New(info, "", 0)
}

// SetErrorOutput set the internal logger.
// Used for package testing.
func (sl *SimpleLogger) SetErrorOutput(error *bytes.Buffer) {
	sl.errorLogger = log.New(error, "", 0)
}
