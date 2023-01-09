package message

import (
	"errors"

	"github.com/william22913/common/error"
)

var UnknownSourceID = error.NewErrorMessage(400, errors.New("BAD_REQUEST"), "Unknown SourceID")
var UnknownMessageID = error.NewErrorMessage(400, errors.New("BAD_REQUEST"), "Unknown MessageID")
