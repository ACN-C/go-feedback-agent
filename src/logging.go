package main // or main, depending on your project

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// This is your global logger
var EventLog *Logging

// Initialize the logger
func SetupLogging() {
	logger := log.New()
	logger.SetOutput(os.Stdout) // safe on Windows

	EventLog = &Logging{Logger: logger}
}

// Optional: set log level safely
func (l *Logging) SetLogLevel(level string) {
	if level == "" {
		level = "info"
	}

	logLevel, err := log.ParseLevel(level)
	if err != nil {
		logLevel = log.InfoLevel
	}

	l.Logger.SetLevel(logLevel)

	if logLevel == log.DebugLevel {
		l.Logger.Infoln("**** DEBUG LOGGING ENABLED ****")
	}
}
