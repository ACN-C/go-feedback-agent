package main

import (
	"os"

	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
)

type Logging struct {
	Logger *log.Logger
}

var eventLog *Logging

func setupLogging() {
	logger := log.New()
	logger.SetOutput(os.Stdout) // Safe on Windows
	eventLog = &Logging{Logger: logger}
}

// program implements service.Interface
type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	// Initialize logging
	if eventLog == nil {
		setupLogging()
	}

	// Do work here
	InitConfig()
	srv := InitServer()

	for {
		conn, err := srv.server.Accept()
		if err != nil {
			eventLog.Logger.Error(err)
			continue
		}
		go handleClient(conn)
	}
}

func (p *program) Stop(s service.Service) error {
	eventLog.Logger.Infoln("Service stopping… reluctantly")
	return nil
}

func main() {
	setupLogging()

	svcConfig := &service.Config{
		Name:        "FeedBackService",
		DisplayName: "TCP Feedback Service",
		Description: "This is a go service to provide system stats",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		eventLog.Logger.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		eventLog.Logger.Error(err)
	}
}
