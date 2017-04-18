package main

import (
	"io/ioutil"
	"os/user"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Consul ConsulConfig `yaml:"consul"`
	Vault  VaultConfig  `yaml:"vault"`
}

type ConsulConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Token     string `yaml:"token"`
	TLS       bool   `yaml:"tls"`
	VerifyTLS bool   `yaml:"verify_tls"`
	KVRoot    string `yaml:"kv_root"`
}

type VaultConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Token     string `yaml:"token"`
	TLS       bool   `yaml:"tls"`
	VerifyTLS bool   `yaml:"verify_tls"`
}

func LoadConfig() error {

	// Configuration default values
	cfg = Config{
		Consul: ConsulConfig{
			Host:      "127.0.0.1",
			Port:      8500,
			Token:     "password",
			TLS:       true,
			VerifyTLS: true,
		},
		Vault: VaultConfig{
			Host:      "127.0.0.1",
			Port:      8200,
			Token:     "password",
			TLS:       true,
			VerifyTLS: true,
		},
	}

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
