package error

func ReformatErrorMessage(err error) DefaultErrorResponse {

	result := DefaultErrorResponse{
		DefaultMessage: DefaultMessage{
			Notification: "error",
		},
	}

	switch errs := err.(type) {
	case ErrorMessages:
		result.Error = DefaultError{
			Code:   errs.Code,
			Reason: err.Error(),
		}
	default:
		result.Error = DefaultError{
			Code:   500,
			Reason: err.Error(),
		}
	}

	return result
}
