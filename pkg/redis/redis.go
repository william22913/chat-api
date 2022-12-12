package redis

import "github.com/go-redis/redis/v8"

type RedisConfiguration struct {
	Address  string `json:"address" envconfig:"address" default:"localhost:6379"`
	Password string `json:"password" envconfig:"password"`
	DB       int    `json:"db" envconfig:"db" default:"0"`
}

func GetRedisConnection(
	config RedisConfiguration,
) *redis.Client {

	opts := &redis.Options{
		Addr: config.Address,
		DB:   config.DB,
	}

	if config.Password != "" {
		opts.Password = config.Password
	}

	return redis.NewClient(opts)

}
