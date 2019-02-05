package main

import (
	"github.com/kelseyhightower/envconfig"
)

const configEnvPrefix = "app"

type DbConfig struct {
	Host     string `envconfig:"host"`
	Port     string `envconfig:"port"`
	Database string `envconfig:"database"`
	User     string `envconfig:"user"`
	Password string `envconfig:"password"`
}

type RpcServerConfig struct {
	Conn string `envconfig:"conn"`
}

type Config struct {
	DB     DbConfig        `envconfig:"db"`
	Server RpcServerConfig `envconfig:"rpcserver"`
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) Init() {
	err := envconfig.Process(configEnvPrefix, cfg)
	if err != nil {
		panic(err)
	}
}
