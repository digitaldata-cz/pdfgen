package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type tConfig struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

func (p *tProgram) loadConfig() {
	p.config = &tConfig{Address: "localhost", Port: "50051"}
	// Try read from env if app is running from container
	if os.Getenv("PS_IP") != "" && os.Getenv("PS_PORT") != "" {
		p.config.Address = os.Getenv("PS_IP")
		p.config.Port = os.Getenv("PS_PORT")
		return
	}
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		logger.Errorf("err-configLoad: %s", err.Error())
		return // If config not found then use default values
	}
	if err := yaml.Unmarshal(file, &p.config); err != nil {
		logger.Errorf("err-configUnmarshal: %s", err.Error())
		os.Exit(1)
	}
}
