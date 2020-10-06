package config

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"

	"go.uber.org/config"
)

//Config is a struct that will be used to parse the enviroment variables located in config.yml
type Config struct {
	Goenv     string
	Directory string
	Exclude   string
	Commands  []string
}

//CfgFactory generates config based on config.yml file
func CfgFactory(configPath string) Config {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	yamlFileReader := bytes.NewReader(yamlFile)

	provider, err := config.NewYAML(config.Source(yamlFileReader))
	if err != nil {
		log.Fatal(err)
	}

	var c Config
	if err := provider.Get("development").Populate(&c); err != nil {
		log.Fatal(err)
	}

	return c
}

// CfgFilePath check if the config.yml directory has been passed
func CfgFilePath() string {
	var configFlag string
	flag.StringVar(&configFlag, "config-file", "config.yml", "a config file path")

	flag.Parse()
	return configFlag
}
