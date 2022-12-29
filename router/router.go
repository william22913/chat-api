package router

import (
	"github.com/rs/zerolog/log"
	wscommunication "github.com/william22913/chat-api/mapping/ws-communication"
	wsmapping "github.com/william22913/chat-api/mapping/ws-mapping"
	"github.com/william22913/chat-api/message"
)

func NewRouter(
	ws wscommunication.Wscommunication,
	f func(message.Message) (
		map[string]wsmapping.WSClientMapping,
		error,
	),
) Router {
	router := &router{
		wsComm: ws,
	}

	router.subscriber = make(chan message.Message)
	router.state = make(chan struct{})

	go router.SubscribeMessage(f)

	return router
}

type router struct {
	subscriber chan message.Message
	state      chan struct{}
	wsComm     wscommunication.Wscommunication
}

func (r *router) UnsubscribeMessage() {
	r.state <- struct{}{}
}

func (r *router) ProcessMessage(msg message.Message) {
	r.subscriber <- msg
}

func (r *router) SubscribeMessage(
	f func(message.Message) (
		map[string]wsmapping.WSClientMapping,
		error,
	),
) {
	for {
		select {
		case msg := <-r.subscriber:
			var err error
			defer func() {
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
			}()

			if msg.Type == message.SendMessage {
				//TODO save message to DB

			} else if msg.Type == message.SendDeliveryStatus {
				//TODO update status to db -Mongo

			} else if msg.Type == message.SendSeen {
				//TODO update status to db -Mongo
			}

			clientMap, err := f(msg)
			if err != nil {
				return
			}

			if clientMap != nil {
				for key := range clientMap {
					r.wsComm.SendMessageWithMapping(clientMap[key], key, msg)
					//TODO Save message to DB -Mongo
				}
			}

		case <-r.state:
			break
		}
	}
}
