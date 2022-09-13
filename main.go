package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/kardianos/service"
)

var (
	svcFlag = flag.String("service", "", "Service controll (start, stop, install, uninstall)")
	logger  service.Logger
)

type tProgram struct {
	exit   chan struct{}
	config *tConfig
}

func init() {
	// Set main function to run on the main thread.
	runtime.LockOSThread()
}

func main() {
	flag.Parse()

	// If running as Windows service current dir can be system32 so must be changed in order to load config.yaml
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(dir)

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	svcConfig := &service.Config{
		Name:        "pdfgen",
		DisplayName: "PDFgen",
		Description: "Service for generating PDF from HTML over gRPC",
		Option:      options,
		EnvVars:     map[string]string{"installed": "true"},
	}

	prg := &tProgram{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			logger.Errorf("Valid actions: %q\n", service.ControlAction)
			logger.Errorf(err.Error())
			os.Exit(1)
		}
		return
	}

	// Start the service.
	if err := s.Run(); err != nil {
		logger.Error(err.Error()) // #nosec G104
		os.Exit(1)
	}
}

func (p *tProgram) Start(s service.Service) error {
	p.exit = make(chan struct{})
	go p.run()
	return nil
}

func (p *tProgram) Stop(s service.Service) error {
	close(p.exit)
	logger.Info("Stopping service")
	// <-time.After(time.Second * 3)
	time.Sleep(3 * time.Second)
	return nil
}
