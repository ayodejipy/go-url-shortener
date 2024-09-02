package logger

import (
	"fmt"
	"log"
	"os"
)

type logLevel int

const (
	INFO logLevel = iota
	WARN
	ERROR
	FATAL
)

var levelStrings = map[logLevel] string{
	INFO: "INFO",
	WARN: "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

type Logger struct {
	infoLogger *log.Logger
	warnLogger *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
	// mu          sync.Mutex
}

func NewLoggerHandler() *Logger {
	return &Logger{
		infoLogger: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLogger: log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		fatalLogger: log.New(os.Stdout, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) log(level logLevel, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)

	switch level {
		case INFO:
			l.infoLogger.Output(2, message)
		case WARN:
			l.warnLogger.Output(2, message)
		case ERROR:
			l.errorLogger.Output(2, message)
		case FATAL:
			l.fatalLogger.Output(2, message)
			os.Exit(1)
	}
}


func (l *Logger) Info(formattedMessage string, value ...interface{}) {
	l.log(INFO, formattedMessage, value...)
}
func (l *Logger) Warn(formattedMessage string, value ...interface{}) {
	l.log(WARN, formattedMessage, value...)
}
func (l *Logger) Error(formattedMessage string, value ...interface{}) {
	l.log(ERROR, formattedMessage, value...)
}
func (l *Logger) Fatal(formattedMessage string, value ...interface{}) {
	l.log(FATAL, formattedMessage, value...)
}