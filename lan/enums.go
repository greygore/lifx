// Code generated by "genlifx"; DO NOT EDIT.
package lan

import "fmt"

// DeviceService is generated from protocol enums
type DeviceService uint8

// DeviceService values
const (
	DeviceServiceUDP = 1
)

func (e DeviceService) String() string {
	switch e {
	case DeviceServiceUDP:
		return "DEVICE_SERVICE_UDP"
	}
	return fmt.Sprintf("DeviceService(%d)", e)
}

// LightWaveform is generated from protocol enums
type LightWaveform uint8

// LightWaveform values
const (
	LightWaveformSaw      = 0
	LightWaveformSine     = 1
	LightWaveformHalfSine = 2
	LightWaveformTriangle = 3
	LightWaveformPulse    = 4
)

func (e LightWaveform) String() string {
	switch e {
	case LightWaveformSaw:
		return "LIGHT_WAVEFORM_SAW"
	case LightWaveformSine:
		return "LIGHT_WAVEFORM_SINE"
	case LightWaveformHalfSine:
		return "LIGHT_WAVEFORM_HALF_SINE"
	case LightWaveformTriangle:
		return "LIGHT_WAVEFORM_TRIANGLE"
	case LightWaveformPulse:
		return "LIGHT_WAVEFORM_PULSE"
	}
	return fmt.Sprintf("LightWaveform(%d)", e)
}

// MultiZoneApplicationRequest is generated from protocol enums
type MultiZoneApplicationRequest uint8

// MultiZoneApplicationRequest values
const (
	MultiZoneApplicationRequestNoApply   = 0
	MultiZoneApplicationRequestApply     = 1
	MultiZoneApplicationRequestApplyOnly = 2
)

func (e MultiZoneApplicationRequest) String() string {
	switch e {
	case MultiZoneApplicationRequestNoApply:
		return "MULTI_ZONE_APPLICATION_REQUEST_NO_APPLY"
	case MultiZoneApplicationRequestApply:
		return "MULTI_ZONE_APPLICATION_REQUEST_APPLY"
	case MultiZoneApplicationRequestApplyOnly:
		return "MULTI_ZONE_APPLICATION_REQUEST_APPLY_ONLY"
	}
	return fmt.Sprintf("MultiZoneApplicationRequest(%d)", e)
}

// MultiZoneEffectType is generated from protocol enums
type MultiZoneEffectType uint8

// MultiZoneEffectType values
const (
	MultiZoneEffectTypeOff  = 0
	MultiZoneEffectTypeMove = 1
)

func (e MultiZoneEffectType) String() string {
	switch e {
	case MultiZoneEffectTypeOff:
		return "MULTI_ZONE_EFFECT_TYPE_OFF"
	case MultiZoneEffectTypeMove:
		return "MULTI_ZONE_EFFECT_TYPE_MOVE"
	}
	return fmt.Sprintf("MultiZoneEffectType(%d)", e)
}

// MultiZoneExtendedApplicationRequest is generated from protocol enums
type MultiZoneExtendedApplicationRequest uint8

// MultiZoneExtendedApplicationRequest values
const (
	MultiZoneExtendedApplicationRequestNoApply   = 0
	MultiZoneExtendedApplicationRequestApply     = 1
	MultiZoneExtendedApplicationRequestApplyOnly = 2
)

func (e MultiZoneExtendedApplicationRequest) String() string {
	switch e {
	case MultiZoneExtendedApplicationRequestNoApply:
		return "MULTI_ZONE_EXTENDED_APPLICATION_REQUEST_NO_APPLY"
	case MultiZoneExtendedApplicationRequestApply:
		return "MULTI_ZONE_EXTENDED_APPLICATION_REQUEST_APPLY"
	case MultiZoneExtendedApplicationRequestApplyOnly:
		return "MULTI_ZONE_EXTENDED_APPLICATION_REQUEST_APPLY_ONLY"
	}
	return fmt.Sprintf("MultiZoneExtendedApplicationRequest(%d)", e)
}

// TileEffectType is generated from protocol enums
type TileEffectType uint8

// TileEffectType values
const (
	TileEffectTypeOff   = 0
	TileEffectTypeMorph = 2
	TileEffectTypeFlame = 3
)

func (e TileEffectType) String() string {
	switch e {
	case TileEffectTypeOff:
		return "TILE_EFFECT_TYPE_OFF"
	case TileEffectTypeMorph:
		return "TILE_EFFECT_TYPE_MORPH"
	case TileEffectTypeFlame:
		return "TILE_EFFECT_TYPE_FLAME"
	}
	return fmt.Sprintf("TileEffectType(%d)", e)
}
