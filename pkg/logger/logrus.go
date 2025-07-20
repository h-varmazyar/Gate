package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"runtime"
)

func NewLogger() *log.Logger {
	logger := log.New()

	logger.SetReportCaller(true)
	logger.SetFormatter(&log.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := fmt.Sprintf("%v: %v(line %v)", frame.File, frame.Function, frame.Line)
			return "", fileName
		},
	})

	return logger
}
