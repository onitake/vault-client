package main

import (
	"io/ioutil"
	"os/user"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	IndexFile string `yaml:"index_file"`
}

func LoadConfig() (error) {

	usr, err := user.Current()
	if err != nil {
		return err
	}

	file, err := ioutil.ReadFile(usr.HomeDir + "/.vaultrc")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return err
	}

	return nil
}