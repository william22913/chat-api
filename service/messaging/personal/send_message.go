package personal

import (
	"net/http"

	"github.com/william22913/chat-api/dto/out"
	"github.com/william22913/chat-api/message"
)

func (pc *personalChatService) SendMessage(
	w http.ResponseWriter,
	r *http.Request,
) {
	action := "personal.service"
	var err error
	var msg message.Message
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

	err = pc.UnmarshalMessage(r, &msg)
	if err != nil {
		return
	}

	err = msg.Validate()
	if err != nil {
		return
	}

	pc.router.ProcessMessage(msg)

	response = out.DefaultResponse{
		Success: true,
		Payload: "success",
	}

}
