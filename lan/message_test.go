package lan_test

import (
	"fmt"
	"testing"

	"github.com/greygore/lifx/lan"
)

func TestExamplePacket(t *testing.T) {
	// From https://lan.developer.lifx.com/docs/building-a-lifx-packet
	expected := "31 00 00 34 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 66 00 00 00 00 55 55 FF FF FF FF AC 0D 00 04 00 00"

	m := lan.Message{
		Payload: &lan.LightSetColor{
			Color: lan.LightHsbk{
				Hue:        21845, // 120° / 360 * 65535 = green
				Saturation: 65535, // max
				Brightness: 65535, // max
				Kelvin:     3500,
			},
			Duration: 1024, // milliseconds
		},
	}
	m.Header.Frame.Tagged = true
	m.Header.Frame.Size = lan.HeaderSize + m.Payload.Size()
	m.Header.ProtocolHeader.Type = lan.LightSetColorType

	b, err := m.MarshalBinary()
	if err != nil {
		t.Fatalf("unable to marshal message to binary: %s", err)
	}
	actual := fmt.Sprintf("% X", b)
	if actual != expected {
		t.Errorf("packet mismatch:\nexpected: %s\nactual:   %s", expected, actual)
	}
}

// TODO(greygore): More tests => different packets, unmarshaling a packet, net.Pipe?
