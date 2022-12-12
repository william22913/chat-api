package auth

import (
	wsmapping "github.com/william22913/chat-api/mapping/ws-mapping"
	"github.com/william22913/chat-api/pkg/service"
)

func NewAuthService(
	wsmapping wsmapping.WSMapping,
) AuthService {
	return &auth{
		wsmapping: wsmapping,
	}
}

type auth struct {
	service.Service

	wsmapping wsmapping.WSMapping
}
