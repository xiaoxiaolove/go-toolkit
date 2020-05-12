package logger

import (
	"github.com/cihub/seelog"
	"log"
)

var (
	logger seelog.LoggerInterface
)

func InitLogger(path string) {
	var err error
	logger, err = seelog.LoggerFromConfigAsFile(path)
	if err != nil {
		log.Fatal("Failed to parse seelog file:", err)
	}
	err = seelog.ReplaceLogger(logger)
	if err != nil {
		log.Fatal("Failed to replace logger:", err)
	}
}


func Trace(v ...interface{}) {
	logger.Trace(v...)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Warn(v ...interface{}) {
	_ = logger.Warn(v...)
}

func Error(v ...interface{}) {
	_ = logger.Error(v...)
}

func Critical(v ...interface{}) {
	_ = logger.Critical(v...)
}

func Tracef(format string, params ...interface{}) {
	logger.Tracef(format, params...)
}

func Debugf(format string, params ...interface{}) {
	logger.Debugf(format, params...)
}

func Infof(format string, params ...interface{}) {
	logger.Infof(format, params...)
}

func Warnf(format string, params ...interface{}) {
	_ = logger.Warnf(format, params...)
}

func Errorf(format string, params ...interface{}) {
	_ = logger.Errorf(format, params...)
}

func Criticalf(format string, params ...interface{}) {
	_ = logger.Criticalf(format, params...)
}
