package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type LoggerService interface {
	Debug(msg string) error
	Info(msg string) error
	Warning(msg string) error
	Error(msg string) error
	StopLog() error
}

type logger struct {
	logger      *log.Logger
	logFile     *os.File
	logFileName string
	logLevel    int
}

var instance *logger
var once sync.Once

func GetLogger() (LoggerService, error) {
	if instance == nil {
		return nil, fmt.Errorf("logger not initiated correctly")
	}
	return instance, nil
}

func (l logger) Debug(msg string) error {
	if l.logger == nil {
		return fmt.Errorf("logger not initiated correctly")
	}
	if 4 <= l.logLevel {
		l.logger.Printf("%v\n", msg)
	}
	return nil
}

func (l logger) Info(msg string) error {
	if l.logger == nil {
		return fmt.Errorf("logger not initiated correctly")
	}
	if 1 <= l.logLevel {
		l.logger.Printf("%v\n", msg)
	}
	return nil
}

func (l logger) Warning(msg string) error {
	if l.logger == nil {
		return fmt.Errorf("logger not initiated correctly")
	}
	if 2 <= l.logLevel {
		l.logger.Printf("%v\n", msg)
	}
	return nil
}

func (l logger) Error(msg string) error {
	if l.logger == nil {
		return fmt.Errorf("logger not initiated correctly")
	}
	if 3 <= l.logLevel {
		l.logger.Printf("%v\n", msg)
	}
	return nil
}

func (l *logger) StopLog() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}

func InitService(fileName string, level int) LoggerService {
	once.Do(func() {
		instance = &logger{}

		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic(err)
		}

		prefix := ""

		switch level {
		case 1:
			prefix = "INFO "
		case 2:
			prefix = "WARNING "
		case 3:
			prefix = "ERROR	"
		case 4:
			prefix = "DEBUG	"
		default:
			panic("log level must be between 1-4, Info, Warning, Error, Debug.")
		}

		instance.logger = log.New(f, prefix, log.Ldate|log.Ltime)
		instance.logFile = f
		instance.logLevel = level
	})
	return instance
}
