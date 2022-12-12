package personal

import (
	"net/http"

	"github.com/william22913/chat-api/dto/out"
	"github.com/william22913/chat-api/messaging"
)

func (pc *personalChatService) SendMessage(
	w http.ResponseWriter,
	r *http.Request,
) {
	action := "personal.service"
	var err error
	var message messaging.Message
	var response interface{}

	defer func() {
		pc.AfterServiceProcess(
			action,
			response,
			w,
			r,
			err,
		)
	}()

	err = pc.UnmarshalMessage(r, &message)
	if err != nil {
		return
	}

	err = message.Validate()
	if err != nil {
		return
	}

	pc.router.ProcessMessage(message)

	response = out.DefaultResponse{
		Success: true,
		Payload: "success",
	}

}
