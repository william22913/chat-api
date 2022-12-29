package personal

import (
	"github.com/william22913/chat-api/pkg/service"
	"github.com/william22913/chat-api/router"
)

func NewPersonalChatService(
	router router.SpecificRouter,
) PersonalChatService {
	return &personalChatService{
		router: router,
	}
}

type personalChatService struct {
	service.Service

	router router.SpecificRouter
}
