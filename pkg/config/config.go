package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Title   string
	Server  Server
	LogInfo LogInfo `toml:"log_info"`
	//Cert    string  `toml:"cert"`
	//Key     string  `toml:"key"`
}

type Server struct {
	Port uint `toml:"port"`
}

type LogInfo struct {
	Level string
}

func NewConfig(buffer []byte) *Config {
	var conf Config
	if err := toml.Unmarshal(buffer, &conf); err != nil {
		log.Fatal(err)
	}
	return &conf
}
