package lan

import (
	"fmt"
	"net"
	"time"
)

// DefaultPort is the default port used by LIFX
const DefaultPort = uint32(56700)

var (
	// BroadcastIP is used to send messages to the entire network
	// Defaults to "255.255.255.255" which should work for most networks
	BroadcastIP = "255.255.255.255"
	// ReadTimeout sets a limit to how long to wait in milliseconds
	// By default (0), reads will block forever
	ReadTimeout time.Duration

	conns map[uint32]net.PacketConn
)

// Conn returns a cached UDP connection for a given port
func Conn(port uint32) (net.PacketConn, error) {
	if conns == nil {
		conns = make(map[uint32]net.PacketConn)
	}

	if _, ok := conns[port]; !ok {
		conn, err := net.ListenPacket("udp", fmt.Sprintf(":%d", port))
		if err != nil {
			return nil, fmt.Errorf("unable to announce on local network address: %s", err)
		}
		conns[port] = conn
	}

	return conns[port], nil
}

// Broadcast sends a message to all devices on the network over the default port
func Broadcast(msg Message) error {
	return Send(BroadcastIP, DefaultPort, msg)
}

// Send will write the message to the specified address and port on the network
func Send(addr string, port uint32, msg Message) error {
	conn, err := Conn(port)
	if err != nil {
		return err
	}

	dst, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return fmt.Errorf("unable to resolve UDP address: %s", err)
	}

	b, err := msg.MarshalBinary()
	if err != nil {
		return err
	}
	n, err := conn.WriteTo(b, dst)
	if err != nil {
		return fmt.Errorf("unable to send: %s", err)
	}
	if n < len(b) {
		return fmt.Errorf("unable to send message: %d bytes of %d sent", n, len(b))
	}

	return nil
}

// Receive attempts to get a single message from the specified port
// If timed out, returns no message and no error
func Receive(port uint32) (host string, msg *Message, err error) {
	buf := make([]byte, largestPacketSize)
	conn, err := Conn(port)
	if err != nil {
		return "", nil, err
	}

	var timeout time.Time
	if ReadTimeout != 0 {
		timeout = time.Now().Add(ReadTimeout)
	}
	if err := conn.SetReadDeadline(timeout); err != nil {
		return "", nil, fmt.Errorf("unable to set read deadline: %s", err)
	}

	n, addr, err := conn.ReadFrom(buf)
	if err != nil {
		if isTimedOut(err) {
			return "", nil, nil
		}
		return "", nil, fmt.Errorf("unable to read message(s): %s", err)
	}
	if n < HeaderSize {
		return "", nil, fmt.Errorf("message size is less than the expected header length of %d", HeaderSize)
	}

	host, _, err = net.SplitHostPort(addr.String())
	if err != nil {
		return "", nil, fmt.Errorf("unable to get host: %s", err)
	}

	var header Header
	if err := header.UnmarshalBinary(buf[:n]); err != nil {
		return "", nil, fmt.Errorf("unable to unmarshal header: %s", err)
	}

	payload, err := PayloadByType(header.ProtocolHeader.Type)
	if err != nil {
		return "", nil, err
	}

	if err := payload.UnmarshalBinary(buf[HeaderSize:n]); err != nil {
		return "", nil, err
	}

	return host, &Message{Header: header, Payload: payload}, nil
}

// timeouter is used to test if an error has timed out
type timeouter interface {
	Timeout() bool
}

func isTimedOut(err error) bool {
	if e, ok := err.(timeouter); ok {
		return e.Timeout()
	}
	return false
}
