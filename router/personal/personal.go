package personal

import (
	"context"

	wscommunication "github.com/william22913/chat-api/mapping/ws-communication"
	wsmapping "github.com/william22913/chat-api/mapping/ws-mapping"
	"github.com/william22913/chat-api/messaging"
	"github.com/william22913/chat-api/router"
)

func NewPersonalChatRouter(
	wsmapping wsmapping.WSMapping,
	wscommunication wscommunication.Wscommunication,
) router.SpecificRouter {

	pcRouter := &personalChatRouter{
		wsmapping: wsmapping,
	}

	pcRouter.router = router.NewRouter(
		wscommunication,
		pcRouter.GetClient,
	)

	return pcRouter
}

type personalChatRouter struct {
	wsmapping wsmapping.WSMapping
	router    router.Router
}

func (pc *personalChatRouter) GetClient(msg messaging.Message) (
	result map[string]wsmapping.WSClientMapping,
	err error,
) {
	ctx := context.Background()
	result = make(map[string]wsmapping.WSClientMapping)

	//TODO Check Client Mapping on db.
	result[msg.DestinationID], err = pc.wsmapping.GetWsClientMapping(ctx, msg.DestinationID)
	return
}

func (pc *personalChatRouter) StopListen() {
	pc.router.UnsubscribeMessage()
}

func (pc *personalChatRouter) ProcessMessage(
	msg messaging.Message,
) {
	pc.router.ProcessMessage(msg)
}
