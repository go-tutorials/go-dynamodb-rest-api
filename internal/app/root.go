package app

import (
	"github.com/common-go/log"
	m "github.com/common-go/middleware"
	"github.com/common-go/dynamodb"
)

type Root struct {
	Server     ServerConfig   `mapstructure:"server"`
	DB         dynamodb.Config `mapstructure:"db"`
	Log        log.Config     `mapstructure:"log"`
	MiddleWare m.LogConfig    `mapstructure:"middleware"`
}

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}
