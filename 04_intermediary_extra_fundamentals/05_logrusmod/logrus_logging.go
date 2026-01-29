package main

import (
	log "github.com/sirupsen/logrus"

	zap "go.uber.org/zap"
)

func main() {
	// create logger
	logger := log.New()

	// set level
	logger.SetLevel(log.InfoLevel)

	// set format
	logger.SetFormatter(&log.JSONFormatter{})

	logger.Info("Hello World")
	logger.Error("Hello World")
	logger.Warn("Hello World")
	logger.Debug("Hello World")

	logger.WithFields(log.Fields{
		"username": "Foo",
		"Method":   "GET",
	}).Info("Request")

	//& logging with zap
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	defer zapLogger.Sync()

	zapLogger.Info("This is a info message", zap.String("username", "Foo"))
}
