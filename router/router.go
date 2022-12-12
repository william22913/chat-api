package router

import (
	"github.com/rs/zerolog/log"
	wscommunication "github.com/william22913/chat-api/mapping/ws-communication"
	wsmapping "github.com/william22913/chat-api/mapping/ws-mapping"
	"github.com/william22913/chat-api/messaging"
)

func NewRouter(
	ws wscommunication.Wscommunication,
	f func(messaging.Message) (
		map[string]wsmapping.WSClientMapping,
		error,
	),
) Router {
	router := &router{
		wsComm: ws,
	}

	router.subscriber = make(chan messaging.Message)
	router.state = make(chan struct{})

	go router.SubscribeMessage(f)

	return router
}

type router struct {
	subscriber chan messaging.Message
	state      chan struct{}
	wsComm     wscommunication.Wscommunication
}

func (r *router) UnsubscribeMessage() {
	r.state <- struct{}{}
}

func (r *router) ProcessMessage(msg messaging.Message) {
	r.subscriber <- msg
}

func (r *router) SubscribeMessage(
	f func(messaging.Message) (
		map[string]wsmapping.WSClientMapping,
		error,
	),
) {
	for {
		select {
		case msg := <-r.subscriber:
			var err error
			if r := recover(); r != nil {
				log.Error().
					Err(r.(error)).
					Str("action", "subscribe.message").
					Caller().
					Send()
			} else {
				if err != nil {
					log.
						Error().
						Caller().
						Err(err).
						Str("action", "subscribe.message").
						Interface("msg", msg).
						Send()
				}
			}

			clientMap, err := f(msg)
			if err != nil {
				return
			}

			if clientMap != nil {
				for key := range clientMap {
					r.wsComm.SendMessageWithMapping(clientMap[key], msg)
				}
			}

		case <-r.state:
			break
		}
	}
}
