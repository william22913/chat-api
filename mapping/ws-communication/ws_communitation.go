package wscommunication

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	wsmapping "github.com/william22913/chat-api/mapping/ws-mapping"
	"github.com/william22913/chat-api/message"
	httpclient "github.com/william22913/common/http_client"
)

func NewWscommunication(
	clientMapping wsmapping.WSMapping,
	httpClient httpclient.APIConnector,
) Wscommunication {
	return &wscommunication{
		clientMapping: clientMapping,
		httpClient:    httpClient,
	}
}

type wscommunication struct {
	clientMapping wsmapping.WSMapping
	httpClient    httpclient.APIConnector
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
	host := fmt.Sprintf("http://%s:8003", ip)
	path := "/message/send"

	wsURL := fmt.Sprintf("%s%s", host, path)
	header := make(map[string]string)
	wsResult := make(map[string]interface{})

	msg.Identity = message.Identity{
		ClientID: client_id,
		Sign:     key,
	}

	httpCode, err := ws.httpClient.HitAPI(
		http.MethodPost,
		host,
		path,
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
