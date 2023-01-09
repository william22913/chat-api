package auth

import (
	"errors"
	"net/http"

	"github.com/william22913/chat-api/message"
)

func (a *auth) Error(
	w http.ResponseWriter,
	r *http.Request,
) {
	action := "error"
	var err error
	var response interface{}

	defer func() {
		a.AfterServiceProcess(
			action,
			response,
			w,
			r,
			err,
		)
	}()

	err = errors.New("test")

}

func (a *auth) Error2(
	w http.ResponseWriter,
	r *http.Request,
) {
	action := "error"
	var err error
	var response interface{}

	defer func() {
		a.AfterServiceProcess(
			action,
			response,
			w,
			r,
			err,
		)
	}()

	err = message.UnknownSourceID

}
