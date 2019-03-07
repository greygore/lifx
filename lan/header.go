package lan

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// HeaderSize is the fixed size of the message header
const HeaderSize = 36

// These are the bits where flags are located
const (
	tagged          = 1 << 13
	acknowledgement = 1 << 1
	response        = 1
)

var endian = binary.LittleEndian

// Header is the non payload portion of the LIFX LAN Protocol packet
// https://lan.developer.lifx.com/v2.0/docs/header-description
type Header struct {
	Frame          Frame
	FrameAddress   FrameAddress
	ProtocolHeader ProtocolHeader
}

// Frame contains the size of the message, LIFX protocol, the usage of the target field (in the FrameAddress), and source identifier
type Frame struct {
	Size        uint16 // 16 - Size of entire message in bytes including this field
	Origin      uint8  //  2 - Message origin indicator: must be zero (0)
	Tagged      bool   //  1 - Determines usage of the Frame Address target field
	Addressable bool   //  1 - Must be 1
	Protocol    uint16 // 12 - Must be 1024 (decimal)
	Source      uint32 // 32 - Source identifier: unique value set by the client, used by responses
}

// FrameAddress contains routing information
type FrameAddress struct {
	Target                  [8]uint8 // 64 - 6 byte device address (MAC address) or zero (0) means all devices. The last two bytes should be 0 bytes.
	_                       [6]uint8 // 48 - Must all be zero (0)
	_                       uint8    //  6 - Reserved
	AcknowledgementRequired bool     //  1 - Acknowledgement message required
	ResponseRequired        bool     //  1 - Response message required
	Sequence                uint8    //  8 - Wrap around message sequence number
}

// ProtocolHeader contains a message type
type ProtocolHeader struct {
	_          uint64 // 64 - Reserved
	Type       uint16 // 16 - Message type determines the payload being used
	reserved16 uint16 // 16 - Reserved
}

// MarshalBinary will convert the header into a proper binary format
// Note that the size is not set and must be done before marshaling
func (h Header) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	otap := uint16(1<<12 + 1024) // Origin, Addressable, and Protocol are fixed
	if h.Frame.Tagged {
		otap = otap | tagged
	}
	var ar uint8
	if h.FrameAddress.AcknowledgementRequired {
		ar = ar | acknowledgement
	}
	if h.FrameAddress.ResponseRequired {
		ar = ar | response
	}

	data := []interface{}{
		h.Frame.Size,
		otap,
		h.Frame.Source,
		h.FrameAddress.Target,
		[6]uint8{}, // reserved
		ar,
		h.FrameAddress.Sequence,
		uint64(0), // reserved
		h.ProtocolHeader.Type,
		uint16(0), // reserved
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

// UnmarshalBinary will convert data read from the network back into a header
func (h *Header) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	var reserved64 uint64
	var otap, reserved16 uint16
	var ar uint8
	vars := []interface{}{
		&h.Frame.Size,
		&otap,
		&h.Frame.Source,
		&h.FrameAddress.Target,
		&[6]uint8{}, // reserved
		&ar,
		&h.FrameAddress.Sequence,
		&reserved64, // reserved
		&h.ProtocolHeader.Type,
		&reserved16, // reserved
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return err
		}
	}
	h.Frame.Tagged = otap&tagged != 0
	h.FrameAddress.AcknowledgementRequired = ar&acknowledgement != 0
	h.FrameAddress.ResponseRequired = ar&response != 0

	if int(h.Frame.Size) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", h.Frame.Size, len(data))
	}

	return nil
}
