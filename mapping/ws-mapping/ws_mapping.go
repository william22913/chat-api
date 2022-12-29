package wsmapping

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

func NewWSMapping(
	redis *redis.Client,
) WSMapping {
	return &wsMapping{
		redis: redis,
	}
}

type wsMapping struct {
	redis *redis.Client
}

func (ws *wsMapping) AddWSClientMapping(
	ctx context.Context,
	clientID string,
	sign string,
	ip string,
) error {
	wsMapping, err := ws.getClientDataFromRedis(ctx, clientID)

	if err != nil {
		return err
	}

	wsMapping[sign] = clientMapping{IP: ip}
	data, err := json.Marshal(wsMapping)

	if err != nil {
		return err
	}

	return ws.redis.Set(ctx, clientID, string(data), 0).Err()
}

func (ws *wsMapping) RemoveWSClientMapping(
	ctx context.Context,
	clientID string,
	sign string,
) error {

	wsMapping, err := ws.getClientDataFromRedis(ctx, clientID)
	if err != nil {
		return err
	}

	delete(wsMapping, sign)
	data, err := json.Marshal(wsMapping)

	if err != nil {
		return err
	}

	return ws.redis.Set(ctx, clientID, string(data), 0).Err()
}

func (ws *wsMapping) GetWsClientMapping(
	ctx context.Context,
	clientID string,
) (
	WSClientMapping,
	error,
) {

	wsMapping, err := ws.getClientDataFromRedis(ctx, clientID)
	if err != nil {
		return nil, err
	}

	return wsMapping, nil
}

func (ws *wsMapping) getClientDataFromRedis(
	ctx context.Context,
	clientID string,
) (
	WSClientMapping,
	error,
) {

	result := make(WSClientMapping)
	redisData, err := ws.redis.Get(ctx, clientID).Result()

	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return result, nil
		}

		return result, err
	}

	_ = json.Unmarshal([]byte(redisData), &result)
	return result, nil
}
