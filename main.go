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
	personalSrv "github.com/william22913/chat-api/service/messaging/personal"
	"github.com/william22913/common/http_client"
	"github.com/william22913/common/metric_middleware"
	"github.com/william22913/common/metrics"
)

func main() {

	config := config.AppConfig
	redis := redis.GetRedisConnection(config.Redis)
	metrics := metrics.NewMetrics()

	httpClient := http_client.NewAPIConnector(metrics)
	metrics_middleware := metric_middleware.NewMetricMiddleware(metrics)

	wsMapping := wsmapping.NewWSMapping(redis)
	wsComm := wscommunication.NewWscommunication(
		wsMapping,
		httpClient,
	)

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
		metrics_middleware,
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
