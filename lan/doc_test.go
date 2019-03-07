package lan_test

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/greygore/lifx/lan"
)

// This example shows a complete workflow for
func Example_completeExample() {
	// Let's broadcast a DeviceGetService message, which will order
	// all the LIFX devices to respond with a DeviceStateService message
	msg := lan.Message{}
	msg.Header.ProtocolHeader.Type = lan.DeviceGetServiceType
	msg.Payload = &lan.DeviceGetService{}
	// Our payload is empty, so the size of the entire packet is
	// the same size as the entire header
	msg.Header.Frame.Size = lan.HeaderSize
	// DeviceGetService requires us to set Tagged to true
	msg.Header.Frame.Tagged = true

	// Broadcast the message to the entire network
	if err := lan.Broadcast(msg); err != nil {
		panic(fmt.Sprintf("Error sending DeviceGetService %s", err))
	}

	// Reads will by default block, so let's set a timeout
	lan.ReadTimeout = 200 * time.Millisecond

	// We need to store some basic info about each device that reports in...
	type Device struct {
		Host string   // Host is the IP address for the bulb
		Port uint32   // Port used to connect to the device
		MAC  [8]uint8 // The MAC address, used in the Target field
	}
	var devices []Device
	for {
		// Pull a message off the network
		host, msg, err := lan.Receive(lan.DefaultPort)
		if err != nil {
			panic(fmt.Sprintf("Error receiving DeviceStateService: %s", err))
		}
		// If message is nil, we timed out, so stop looking for messages
		if msg == nil {
			break
		}
		// Skip non-DeviceStateService messages
		if msg.Header.ProtocolHeader.Type != lan.DeviceStateServiceType {
			continue
		}
		// Assert our payload to the appropriate type to access its fields
		payload := msg.Payload.(*lan.DeviceStateService)
		// Skip non-UDP service messages
		if payload.Service != lan.DeviceServiceUDP {
			continue
		}
		// Add our device to the list
		devices = append(devices, Device{Host: host, Port: payload.Port, MAC: msg.Header.FrameAddress.Target})
	}

	// Now let's get a label for each device using DeviceGetLabel
	for _, device := range devices {
		msg := lan.Message{}
		msg.Header.ProtocolHeader.Type = lan.DeviceGetLabelType
		msg.Payload = &lan.DeviceGetLabel{}
		// DeviceGetLabel is also an empty payload...
		msg.Header.Frame.Size = lan.HeaderSize
		// Because we are contacting a specific device, we need to include
		// the MAC address in the Target
		msg.Header.FrameAddress.Target = device.MAC

		// Now we can send our message to the device using its stored address & port
		if err := lan.Send(device.Host, device.Port, msg); err != nil {
			log.Fatalf("Error sending Device GetLabel %s", err)
		}

		// Because we might receive other messages while looking for
		// DeviceStateService, we loop until we find one
		for {
			// We need to listen on the port that the device requested we use
			_, m, err := lan.Receive(device.Port)
			if err != nil {
				log.Fatalf("Error receiving Device StateLabel: %s", err)
			}
			// If the read timed out, the bulb may not have replied yet
			if m == nil {
				continue
			}
			// Skip any non DeviceState messages
			if m.Header.ProtocolHeader.Type != lan.DeviceStateLabelType {
				continue
			}

			payload := m.Payload.(*lan.DeviceStateLabel)
			// Our label is in a fixed array of bytes, so trim a slice to get
			// a string of the right length
			label := strings.TrimRight(string(payload.Label[:]), " ")

			// Print our label and break out of the loop
			fmt.Printf("Device: %s\n", label)
			break
		}
	}
}
