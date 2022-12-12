package router

import (
	wsmapping "github.com/william22913/chat-api/mapping/ws-mapping"
	"github.com/william22913/chat-api/messaging"
)

type SpecificRouter interface {
	GetClient(messaging.Message) (
		map[string]wsmapping.WSClientMapping,
		error,
	)

	ProcessMessage(messaging.Message)

	StopListen()
}

type Router interface {
	SubscribeMessage(
		f func(messaging.Message) (
			map[string]wsmapping.WSClientMapping,
			error,
		),
	)

	UnsubscribeMessage()

	ProcessMessage(messaging.Message)
}
