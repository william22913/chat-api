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

	Error(
		w http.ResponseWriter,
		r *http.Request,
	)

	Error2(
		w http.ResponseWriter,
		r *http.Request,
	)
}
