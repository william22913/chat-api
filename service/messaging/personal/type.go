package personal

import "net/http"

type PersonalChatService interface {
	SendMessage(
		w http.ResponseWriter,
		r *http.Request,
	)
}
