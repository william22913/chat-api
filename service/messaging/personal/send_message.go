package personal

import (
	"net/http"

	"github.com/william22913/chat-api/message"
	"github.com/william22913/common/dto/out"
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
