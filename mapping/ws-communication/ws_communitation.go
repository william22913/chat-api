package wscommunication

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	wsmapping "github.com/william22913/chat-api/mapping/ws-mapping"
	"github.com/william22913/chat-api/messaging"
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

func (ws *wscommunication) SendMessageToClient(
	ctx context.Context,
	message messaging.Message,
) {
	mapping, err := ws.clientMapping.GetWsClientMapping(ctx, message.SourceID)
	if err != nil {
		log.Error().
			Caller().
			Interface("message", message).
			Str("client", message.SourceID).
			Err(err).
			Msg("error when get ws client mapping")
	}

	ws.SendMessageWithMapping(mapping, message)
}

func (ws *wscommunication) SendMessageWithMapping(
	mapping wsmapping.WSClientMapping,
	message messaging.Message,
) {
	for key := range mapping {
		wsConn := mapping[key]
		ws.sendToWS(
			wsConn.IP,
			key,
			message,
		)
	}
}

func (ws *wscommunication) sendToWS(
	ip string,
	key string,
	message messaging.Message,
) {
	wsURL := fmt.Sprintf("http://%s:8001/message/send", ip)
	header := make(map[string]string)
	wsResult := make(map[string]interface{})

	keys := strings.Split(key, ".")
	if len(keys) == 2 {
		message.Receiver = messaging.Receiver{
			Device: keys[0],
			Key:    keys[1],
		}
	}

	httpCode, err := httpclient.HitAPI(
		http.MethodPost,
		wsURL,
		header,
		message,
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
