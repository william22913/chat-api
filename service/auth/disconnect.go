package auth

import (
	"context"
	"net/http"

	"github.com/william22913/chat-api/dto/in"
	"github.com/william22913/chat-api/dto/out"
)

func (a *auth) ClientDisconnect(
	w http.ResponseWriter,
	r *http.Request,
) {
	action := "disconnect"
	var err error
	var connectDTO in.ConnectDTO
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

	err = a.UnmarshalMessage(r, &connectDTO)
	if err != nil {
		return
	}

	err = connectDTO.Validate()
	if err != nil {
		return
	}

	err = a.serveDisconnect(connectDTO)
	if err != nil {
		return
	}

	response = out.DefaultResponse{
		Success: true,
		Payload: "success",
	}

}

func (a *auth) serveDisconnect(
	connectDTO in.ConnectDTO,
) error {
	ctx := context.Background()

	return a.wsmapping.RemoveWSClientMapping(
		ctx,
		connectDTO.ClientID,
		connectDTO.Sign,
	)
}
