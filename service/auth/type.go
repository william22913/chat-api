package auth

import (
	"net/http"
)

type AuthService interface {
	ClientConnect(
		w http.ResponseWriter,
		r *http.Request,
	)

	ClientDisconnect(
		w http.ResponseWriter,
		r *http.Request,
	)
}
