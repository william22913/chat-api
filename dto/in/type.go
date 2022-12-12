package in

import (
	"errors"

	srvError "github.com/william22913/chat-api/pkg/error"
)

var ErrInvalidMessage = srvError.NewErrorMessage(400, errors.New("BAD_REQUEST"), "Invalid Message")
