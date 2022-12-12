package wsmapping

import "context"

type WSClientMapping map[string]clientMapping

type clientMapping struct {
	IP string `json:"ip"`
}

type WSMapping interface {
	AddWSClientMapping(
		ctx context.Context,
		clientID string,
		device string,
		key string,
		ip string,
	) error

	RemoveWSClientMapping(
		ctx context.Context,
		clientID string,
		device string,
		key string,
	) error

	GetWsClientMapping(
		ctx context.Context,
		clientID string,
	) (
		WSClientMapping,
		error,
	)
}
