package in

type ConnectDTO struct {
	ClientID string `json:"client_id"`
	SocketIP string `json:"socket_ip"`
	Sign     string `json:"sign"`
}

func (c *ConnectDTO) Validate() error {
	if c.ClientID == "" {
		return ErrInvalidMessage
	}

	return nil
}
