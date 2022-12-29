package router

import (
	wsmapping "github.com/william22913/chat-api/mapping/ws-mapping"
	"github.com/william22913/chat-api/message"
)

type SpecificRouter interface {
	GetClient(message.Message) (
		map[string]wsmapping.WSClientMapping,
		error,
	)

	ProcessMessage(message.Message)

	StopListen()
}

type Router interface {
	SubscribeMessage(
		f func(message.Message) (
			map[string]wsmapping.WSClientMapping,
			error,
		),
	)

	UnsubscribeMessage()

	ProcessMessage(message.Message)
}
