package config

import (
	"encoding/json"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/william22913/chat-api/pkg/redis"
)

var (
	AppConfig Configuration
)

type http struct {
	Port int `envconfig:"port"`
}

type logs struct {
	Level string `envconfig:"level" default:"info"`
}

type Configuration struct {
	Http  http                     `envconfig:"http"`
	Redis redis.RedisConfiguration `envconfig:"redis"`
	Log   logs                     `envconfig:"log"`
}

func init() {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if err := envconfig.Process(
		"chat_api",
		&AppConfig,
	); err != nil {
		log.Fatal().
			Str("action", "parse.conf").
			Err(err).
			Msg("parse configuration fails")
	}

	PrintConfig(AppConfig)
}

func PrintConfig(c Configuration) {
	data, _ := json.MarshalIndent(c, "", "\t")
	fmt.Println(string(data))
}
