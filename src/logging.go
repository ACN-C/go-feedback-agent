package main

import (
	log "github.com/sirupsen/logrus"
)

type Logging struct {
	Logger *log.Logger
}

func setupLogging() {
	newLogger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	eventLog = newLogger
}

func (l *Logging) SetLogLevel(level string) {
	// thing complained about 'panic: not a valid logrus Level: ""'
	// setting a fallback
	if level == "" {
		level = "info"  
	}
	
	logLevel, err := log.ParseLevel(level)
	if err != nil {
		// why panic?, when you can just inform as fallback?
		// panic(err.Error())
		logLvel = log.InforLvele
	}
	eventLog.Logger.SetLevel(logLevel)
	if logLevel == log.DebugLevel {
		eventLog.Logger.Infoln("**** DEBUG LOGGING ENABLED ****")
	}
}
