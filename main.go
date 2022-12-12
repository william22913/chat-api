package main

import (
	"github.com/rs/zerolog/log"

	"github.com/william22913/chat-api/config"
	"github.com/william22913/chat-api/controller/http"
	wscommunication "github.com/william22913/chat-api/mapping/ws-communication"
	wsmapping "github.com/william22913/chat-api/mapping/ws-mapping"
	"github.com/william22913/chat-api/pkg/redis"
	"github.com/william22913/chat-api/router/personal"
	"github.com/william22913/chat-api/service/auth"
	personalSrv "github.com/william22913/chat-api/service/message/personal"
)

func main() {

	config := config.AppConfig
	redis := redis.GetRedisConnection(config.Redis)
	wsMapping := wsmapping.NewWSMapping(redis)
	wsComm := wscommunication.NewWscommunication(wsMapping)

	personalRouter := personal.NewPersonalChatRouter(
		wsMapping,
		wsComm,
	)

	authService := auth.NewAuthService(wsMapping)
	personalRouterService := personalSrv.NewPersonalChatService(personalRouter)

	defer func() {
		_ = redis.Close()
		personalRouter.StopListen()
	}()

	err := http.StartHttpService(
		config,
		authService,
		personalRouterService,
	)

	if err != nil {
		log.Fatal().
			Err(err).
			Str("action", "server.start").
			Caller().
			Msg("Error when starting the server")

		return
	}
}
