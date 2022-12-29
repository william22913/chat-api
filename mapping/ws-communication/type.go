package wscommunication

import (
	wsmapping "github.com/william22913/chat-api/mapping/ws-mapping"
	"github.com/william22913/chat-api/message"
)

type Wscommunication interface {
	SendMessageWithMapping(
		mapping wsmapping.WSClientMapping,
		client_id string,
		message message.Message,
	)
}
