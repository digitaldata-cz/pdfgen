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
	if os.Getenv("IP") != "" && os.Getenv("PORT") != "" {
		p.config.Address = os.Getenv("IP")
		p.config.Port = os.Getenv("PORT")
		return
	}
	if _, err := os.Stat("config.yaml"); err == nil {
		// Read config from file
		data, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			logger.Error(err.Error())
			return
		}
		err = yaml.Unmarshal(data, p.config)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
		return
	}
}
