package app

import (
	"github.com/core-go/dynamodb"
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
)

type Root struct {
	Server     ServerConfig    `mapstructure:"server"`
	DB         dynamodb.Config `mapstructure:"db"`
	Log        log.Config      `mapstructure:"log"`
	MiddleWare mid.LogConfig   `mapstructure:"middleware"`
}

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int64  `mapstructure:"port"`
}
