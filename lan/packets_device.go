// Code generated by "genlifx"; DO NOT EDIT.
package lan

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// message types
const (
	DeviceAcknowledgementType   = 45
	DeviceEchoRequestType       = 58
	DeviceEchoResponseType      = 59
	DeviceGetGroupType          = 51
	DeviceGetHostFirmwareType   = 14
	DeviceGetHostInfoType       = 12
	DeviceGetInfoType           = 34
	DeviceGetLabelType          = 23
	DeviceGetLocationType       = 48
	DeviceGetPowerType          = 20
	DeviceGetServiceType        = 2
	DeviceGetVersionType        = 32
	DeviceGetWifiFirmwareType   = 18
	DeviceGetWifiInfoType       = 16
	DeviceSetGroupType          = 52
	DeviceSetLabelType          = 24
	DeviceSetLocationType       = 49
	DeviceSetPowerType          = 21
	DeviceStateGroupType        = 53
	DeviceStateHostFirmwareType = 15
	DeviceStateHostInfoType     = 13
	DeviceStateInfoType         = 35
	DeviceStateLabelType        = 25
	DeviceStateLocationType     = 50
	DeviceStatePowerType        = 22
	DeviceStateServiceType      = 3
	DeviceStateVersionType      = 33
	DeviceStateWifiFirmwareType = 19
	DeviceStateWifiInfoType     = 17
)

/////////////////////////////////////////////////////////////////////////////

type DeviceAcknowledgement struct{}

func (m DeviceAcknowledgement) Size() uint16 {
	return 0
}

func (m DeviceAcknowledgement) Type() uint16 {
	return DeviceAcknowledgementType
}

func (m DeviceAcknowledgement) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (m *DeviceAcknowledgement) UnmarshalBinary(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceEchoRequest struct {
	Payload [64]byte
}

func (m DeviceEchoRequest) Size() uint16 {
	return 64
}

func (m DeviceEchoRequest) Type() uint16 {
	return DeviceEchoRequestType
}

func (m DeviceEchoRequest) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Payload,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceEchoRequest) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Payload,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceEchoResponse struct {
	Payload [64]byte
}

func (m DeviceEchoResponse) Size() uint16 {
	return 64
}

func (m DeviceEchoResponse) Type() uint16 {
	return DeviceEchoResponseType
}

func (m DeviceEchoResponse) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Payload,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceEchoResponse) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Payload,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceGetGroup struct{}

func (m DeviceGetGroup) Size() uint16 {
	return 0
}

func (m DeviceGetGroup) Type() uint16 {
	return DeviceGetGroupType
}

func (m DeviceGetGroup) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (m *DeviceGetGroup) UnmarshalBinary(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceGetHostFirmware struct{}

func (m DeviceGetHostFirmware) Size() uint16 {
	return 0
}

func (m DeviceGetHostFirmware) Type() uint16 {
	return DeviceGetHostFirmwareType
}

func (m DeviceGetHostFirmware) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (m *DeviceGetHostFirmware) UnmarshalBinary(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceGetHostInfo struct{}

func (m DeviceGetHostInfo) Size() uint16 {
	return 0
}

func (m DeviceGetHostInfo) Type() uint16 {
	return DeviceGetHostInfoType
}

func (m DeviceGetHostInfo) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (m *DeviceGetHostInfo) UnmarshalBinary(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceGetInfo struct{}

func (m DeviceGetInfo) Size() uint16 {
	return 0
}

func (m DeviceGetInfo) Type() uint16 {
	return DeviceGetInfoType
}

func (m DeviceGetInfo) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (m *DeviceGetInfo) UnmarshalBinary(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceGetLabel struct{}

func (m DeviceGetLabel) Size() uint16 {
	return 0
}

func (m DeviceGetLabel) Type() uint16 {
	return DeviceGetLabelType
}

func (m DeviceGetLabel) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (m *DeviceGetLabel) UnmarshalBinary(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceGetLocation struct{}

func (m DeviceGetLocation) Size() uint16 {
	return 0
}

func (m DeviceGetLocation) Type() uint16 {
	return DeviceGetLocationType
}

func (m DeviceGetLocation) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (m *DeviceGetLocation) UnmarshalBinary(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceGetPower struct{}

func (m DeviceGetPower) Size() uint16 {
	return 0
}

func (m DeviceGetPower) Type() uint16 {
	return DeviceGetPowerType
}

func (m DeviceGetPower) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (m *DeviceGetPower) UnmarshalBinary(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceGetService struct{}

func (m DeviceGetService) Size() uint16 {
	return 0
}

func (m DeviceGetService) Type() uint16 {
	return DeviceGetServiceType
}

func (m DeviceGetService) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (m *DeviceGetService) UnmarshalBinary(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceGetVersion struct{}

func (m DeviceGetVersion) Size() uint16 {
	return 0
}

func (m DeviceGetVersion) Type() uint16 {
	return DeviceGetVersionType
}

func (m DeviceGetVersion) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (m *DeviceGetVersion) UnmarshalBinary(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceGetWifiFirmware struct{}

func (m DeviceGetWifiFirmware) Size() uint16 {
	return 0
}

func (m DeviceGetWifiFirmware) Type() uint16 {
	return DeviceGetWifiFirmwareType
}

func (m DeviceGetWifiFirmware) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (m *DeviceGetWifiFirmware) UnmarshalBinary(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceGetWifiInfo struct{}

func (m DeviceGetWifiInfo) Size() uint16 {
	return 0
}

func (m DeviceGetWifiInfo) Type() uint16 {
	return DeviceGetWifiInfoType
}

func (m DeviceGetWifiInfo) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (m *DeviceGetWifiInfo) UnmarshalBinary(data []byte) error {
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceSetGroup struct {
	Group     [16]byte
	Label     [32]byte
	UpdatedAt uint64
}

func (m DeviceSetGroup) Size() uint16 {
	return 56
}

func (m DeviceSetGroup) Type() uint16 {
	return DeviceSetGroupType
}

func (m DeviceSetGroup) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Group,
		m.Label,
		m.UpdatedAt,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceSetGroup) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Group,
		&m.Label,
		&m.UpdatedAt,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceSetLabel struct {
	Label [32]byte
}

func (m DeviceSetLabel) Size() uint16 {
	return 32
}

func (m DeviceSetLabel) Type() uint16 {
	return DeviceSetLabelType
}

func (m DeviceSetLabel) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Label,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceSetLabel) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Label,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceSetLocation struct {
	Location  [16]byte
	Label     [32]byte
	UpdatedAt uint64
}

func (m DeviceSetLocation) Size() uint16 {
	return 56
}

func (m DeviceSetLocation) Type() uint16 {
	return DeviceSetLocationType
}

func (m DeviceSetLocation) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Location,
		m.Label,
		m.UpdatedAt,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceSetLocation) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Location,
		&m.Label,
		&m.UpdatedAt,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceSetPower struct {
	Level uint16
}

func (m DeviceSetPower) Size() uint16 {
	return 2
}

func (m DeviceSetPower) Type() uint16 {
	return DeviceSetPowerType
}

func (m DeviceSetPower) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Level,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceSetPower) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Level,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceStateGroup struct {
	Group     [16]byte
	Label     [32]byte
	UpdatedAt uint64
}

func (m DeviceStateGroup) Size() uint16 {
	return 56
}

func (m DeviceStateGroup) Type() uint16 {
	return DeviceStateGroupType
}

func (m DeviceStateGroup) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Group,
		m.Label,
		m.UpdatedAt,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceStateGroup) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Group,
		&m.Label,
		&m.UpdatedAt,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceStateHostFirmware struct {
	Build     uint64
	Reserved1 [8]byte
	Version   uint32
}

func (m DeviceStateHostFirmware) Size() uint16 {
	return 20
}

func (m DeviceStateHostFirmware) Type() uint16 {
	return DeviceStateHostFirmwareType
}

func (m DeviceStateHostFirmware) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Build,
		m.Reserved1,
		m.Version,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceStateHostFirmware) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Build,
		&m.Reserved1,
		&m.Version,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceStateHostInfo struct {
	Signal    float32
	Tx        uint32
	Rx        uint32
	Reserved3 [2]byte
}

func (m DeviceStateHostInfo) Size() uint16 {
	return 14
}

func (m DeviceStateHostInfo) Type() uint16 {
	return DeviceStateHostInfoType
}

func (m DeviceStateHostInfo) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Signal,
		m.Tx,
		m.Rx,
		m.Reserved3,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceStateHostInfo) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Signal,
		&m.Tx,
		&m.Rx,
		&m.Reserved3,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceStateInfo struct {
	Time     uint64
	Uptime   uint64
	Downtime uint64
}

func (m DeviceStateInfo) Size() uint16 {
	return 24
}

func (m DeviceStateInfo) Type() uint16 {
	return DeviceStateInfoType
}

func (m DeviceStateInfo) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Time,
		m.Uptime,
		m.Downtime,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceStateInfo) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Time,
		&m.Uptime,
		&m.Downtime,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceStateLabel struct {
	Label [32]byte
}

func (m DeviceStateLabel) Size() uint16 {
	return 32
}

func (m DeviceStateLabel) Type() uint16 {
	return DeviceStateLabelType
}

func (m DeviceStateLabel) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Label,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceStateLabel) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Label,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceStateLocation struct {
	Location  [16]byte
	Label     [32]byte
	UpdatedAt uint64
}

func (m DeviceStateLocation) Size() uint16 {
	return 56
}

func (m DeviceStateLocation) Type() uint16 {
	return DeviceStateLocationType
}

func (m DeviceStateLocation) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Location,
		m.Label,
		m.UpdatedAt,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceStateLocation) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Location,
		&m.Label,
		&m.UpdatedAt,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceStatePower struct {
	Level uint16
}

func (m DeviceStatePower) Size() uint16 {
	return 2
}

func (m DeviceStatePower) Type() uint16 {
	return DeviceStatePowerType
}

func (m DeviceStatePower) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Level,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceStatePower) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Level,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceStateService struct {
	Service DeviceService
	Port    uint32
}

func (m DeviceStateService) Size() uint16 {
	return 5
}

func (m DeviceStateService) Type() uint16 {
	return DeviceStateServiceType
}

func (m DeviceStateService) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Service,
		m.Port,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceStateService) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Service,
		&m.Port,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceStateVersion struct {
	Vendor  uint32
	Product uint32
	Version uint32
}

func (m DeviceStateVersion) Size() uint16 {
	return 12
}

func (m DeviceStateVersion) Type() uint16 {
	return DeviceStateVersionType
}

func (m DeviceStateVersion) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Vendor,
		m.Product,
		m.Version,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceStateVersion) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Vendor,
		&m.Product,
		&m.Version,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceStateWifiFirmware struct {
	Build     uint64
	Reserved1 [8]byte
	Version   uint32
}

func (m DeviceStateWifiFirmware) Size() uint16 {
	return 20
}

func (m DeviceStateWifiFirmware) Type() uint16 {
	return DeviceStateWifiFirmwareType
}

func (m DeviceStateWifiFirmware) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Build,
		m.Reserved1,
		m.Version,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceStateWifiFirmware) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Build,
		&m.Reserved1,
		&m.Version,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////

type DeviceStateWifiInfo struct {
	Signal    float32
	Tx        uint32
	Rx        uint32
	Reserved3 [2]byte
}

func (m DeviceStateWifiInfo) Size() uint16 {
	return 14
}

func (m DeviceStateWifiInfo) Type() uint16 {
	return DeviceStateWifiInfoType
}

func (m DeviceStateWifiInfo) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	data := []interface{}{
		m.Signal,
		m.Tx,
		m.Rx,
		m.Reserved3,
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func (m *DeviceStateWifiInfo) UnmarshalBinary(data []byte) error {
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		&m.Signal,
		&m.Tx,
		&m.Rx,
		&m.Reserved3,
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}

	return nil
}