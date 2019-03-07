package lan

import "fmt"

// Payloader defines how a payload can be handled by the Message
type Payloader interface {
	Size() uint16
	Type() uint16
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
}

// Message is a LIFX LAN Protocol packet
type Message struct {
	Header  Header
	Payload Payloader
}

// MarshalBinary encodes the message into binary
func (m Message) MarshalBinary() ([]byte, error) {
	header, err := m.Header.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("unable to marshal header: %s", err)
	}

	payload, err := m.Payload.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("unable to marshal payload: %s", err)
	}

	return append(header, payload...), nil
}
