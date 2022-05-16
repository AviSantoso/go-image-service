package logger

import (
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	Id   string
	Path string
}

func New(path string, id string) Logger {
	log.SetFormatter(&log.JSONFormatter{})
	logger := Logger{Id: id, Path: path}
	return logger
}

func (logger Logger) Info(message string) {
	log.WithFields(log.Fields{
		"path": logger.Path,
		"id":   logger.Id,
	}).Info(message)
}
