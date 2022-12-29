package wscommunication

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	wsmapping "github.com/william22913/chat-api/mapping/ws-mapping"
	"github.com/william22913/chat-api/message"
	httpclient "github.com/william22913/chat-api/pkg/http_client"
)

func NewWscommunication(
	clientMapping wsmapping.WSMapping,
) Wscommunication {
	return &wscommunication{
		clientMapping: clientMapping,
	}
}

type wscommunication struct {
	clientMapping wsmapping.WSMapping
}

func (ws *wscommunication) SendMessageWithMapping(
	mapping wsmapping.WSClientMapping,
	client_id string,
	message message.Message,
) {
	for key := range mapping {
		wsConn := mapping[key]
		ws.sendToWS(
			wsConn.IP,
			key,
			client_id,
			message,
		)
	}
}

func (ws *wscommunication) sendToWS(
	ip string,
	key string,
	client_id string,
	msg message.Message,
) {
	wsURL := fmt.Sprintf("http://%s:8003/message/send", ip)
	header := make(map[string]string)
	wsResult := make(map[string]interface{})

	msg.Identity = message.Identity{
		ClientID: client_id,
		Sign:     key,
	}

	httpCode, err := httpclient.HitAPI(
		http.MethodPost,
		wsURL,
		header,
		msg,
		&wsResult,
	)

	if err != nil {
		log.
			Error().
			Caller().
			Err(err).
			Str("url", wsURL).
			Str("action", "hit.api").
			Send()
		return
	}

	if httpCode != 200 {
		log.
			Info().
			Caller().
			Str("url", wsURL).
			Interface("result", wsResult).
			Str("action", "hit.api").
			Send()
	}

}
