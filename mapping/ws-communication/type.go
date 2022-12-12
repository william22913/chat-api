package wscommunication

import (
	"context"

	wsmapping "github.com/william22913/chat-api/mapping/ws-mapping"
	"github.com/william22913/chat-api/messaging"
)

type Wscommunication interface {
	SendMessageToClient(
		ctx context.Context,
		message messaging.Message,
	)

	SendMessageWithMapping(
		mapping wsmapping.WSClientMapping,
		message messaging.Message,
	)
}
