package error

import "errors"

func NewErrorMessage(code int, err error, detail string) error {
	return ErrorMessages{
		Code:         code,
		ErrorMessage: err,
		ErrorDetail:  detail,
	}
}

type ErrorMessages struct {
	Code         int
	ErrorMessage error
	ErrorDetail  string
}

func (e ErrorMessages) Error() string {
	return e.ErrorMessage.Error()
}

var ErrUnauthorized = NewErrorMessage(401, errors.New("UNAUTHORIZED_CLIENT"), "")
var ErrBadRequest = NewErrorMessage(400, errors.New("BAD_REQUEST"), "")
