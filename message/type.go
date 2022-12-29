package message

import (
	"time"
)

type MessageType string

const (
	SendMessage        MessageType = "send"
	SendDeliveryStatus MessageType = "delivered"
	SendSeen           MessageType = "seen"
	Other              MessageType = "other"
)

type Message struct {
	MessageID     string      `json:"message_id"`
	SourceID      string      `json:"source_id"`
	DestinationID string      `json:"destination_id"`
	Guarantee     bool        `json:"guarantee"`
	Message       string      `json:"message"`
	Order         interface{} `json:"order"`
	Time          time.Time   `json:"time"`
	Type          MessageType `json:"type"`
	Identity      Identity    `json:"identity"`
}

type Identity struct {
	ClientID string `json:"client_id"`
	Sign     string `json:"sign"`
}

func (m *Message) Validate() error {
	//TODO Validate Message

	if m.SourceID == "" {
		return UnknownSourceID
	}

	err := m.validateMessageType()
	if err != nil {
		return err
	}

	return nil

}

func (m *Message) validateMessageType() error {

	if m.Type == SendMessage ||
		m.Type == SendDeliveryStatus ||
		m.Type == SendSeen {
		m.Guarantee = true
	}

	if m.Type == SendDeliveryStatus ||
		m.Type == SendSeen {
		if m.MessageID == "" {
			return UnknownMessageID
		}
	}

	return nil
}
