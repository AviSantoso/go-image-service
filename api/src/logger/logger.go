package logger

import (
	"io"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	Id   string
	Path string
}

func New(writer io.Writer, path string, id string) Logger {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(writer)
	logger := Logger{Id: id, Path: path}
	return logger
}

func (logger Logger) Info(message string) *log.Entry {
	out := log.WithFields(log.Fields{
		"path": logger.Path,
		"id":   logger.Id,
	})
	out.Info(message)
	return out
}

func (logger Logger) Error(message string) *log.Entry {
	out := log.WithFields(log.Fields{
		"path": logger.Path,
		"id":   logger.Id,
	})
	out.Error(message)
	return out
}
