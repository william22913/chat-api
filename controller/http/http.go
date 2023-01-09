package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"github.com/william22913/chat-api/config"
	"github.com/william22913/chat-api/service/auth"
	"github.com/william22913/chat-api/service/messaging/personal"
	"github.com/william22913/common/metric_middleware"
)

func StartHttpService(
	config config.Configuration,
	authService auth.AuthService,
	personalRouterService personal.PersonalChatService,
	metric_middleware metric_middleware.MetricMiddleware,
) error {

	router := mux.NewRouter()

	router.HandleFunc("/metrics", promhttp.Handler().ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/auth/connect", authService.ClientConnect).Methods(http.MethodPost)
	router.HandleFunc("/auth/disconnect", authService.ClientDisconnect).Methods(http.MethodPost)
	router.HandleFunc("/auth/error", authService.Error).Methods(http.MethodGet)
	router.HandleFunc("/auth/error2", authService.Error2).Methods(http.MethodGet)
	router.HandleFunc("/message/personal", personalRouterService.SendMessage).Methods(http.MethodPost)

	router.Use(middleware)
	router.Use(metric_middleware.Serve)

	log.Info().
		Str("action", "server.start").
		Int("port", config.Http.Port).
		Msg("HTTP Server Start.")

	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.Http.Port), router)

}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI != "/metrics" && r.RequestURI != "/health" {
			log.Info().
				Str("action", "middleware.api.call").
				Str("url", r.URL.Path).
				Str("method", r.Method).
				Msg("Api Called.")
		}

		h.ServeHTTP(w, r)
	})
}
