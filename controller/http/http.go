package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/william22913/chat-api/config"
	"github.com/william22913/chat-api/service/auth"
	"github.com/william22913/chat-api/service/messaging/personal"
)

func StartHttpService(
	config config.Configuration,
	authService auth.AuthService,
	personalRouterService personal.PersonalChatService,
) error {

	router := mux.NewRouter()

	router.HandleFunc("/auth/connect", authService.ClientConnect).Methods(http.MethodPost)
	router.HandleFunc("/auth/disconnect", authService.ClientDisconnect).Methods(http.MethodPost)
	router.HandleFunc("/message/personal", personalRouterService.SendMessage).Methods(http.MethodPost)

	router.Use(middleware)

	log.Info().
		Str("action", "server.start").
		Int("port", config.Http.Port).
		Msg("HTTP Server Start.")

	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.Http.Port), router)

}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Info().
			Str("action", "middleware.api.call").
			Str("url", r.URL.Path).
			Str("method", r.Method).
			Msg("Api Called.")

		h.ServeHTTP(w, r)
	})
}
