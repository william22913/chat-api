package in

import (
	"errors"

	srvError "github.com/william22913/common/error"
)

var ErrInvalidMessage = srvError.NewErrorMessage(400, errors.New("BAD_REQUEST"), "Invalid Message")
