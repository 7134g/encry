package config

import (
	"encry/logs"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Setting struct {
	Address struct {
		Remote string `yaml:"remote"`
		Local  string `yaml:"local"`
	}
	Other struct {
		LogStatus string `yaml:"logStatus"`
	}
}

func LoadYaml() {
	conf := new(Setting)
	yamlFile, err := ioutil.ReadFile("env.yaml")
	if err != nil {
		logs.Exit("The y", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		logs.Error("Unmarshal: %v", err)
	}
}
