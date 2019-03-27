// Code generated by "genlifx"; DO NOT EDIT.
package lan

import "fmt"

const largestPacketSize = 882

func PayloadByType(payloadType uint16) (Payloader, error) {
	switch payloadType {
	case DeviceAcknowledgementType:
		return &DeviceAcknowledgement{}, nil
	case DeviceEchoRequestType:
		return &DeviceEchoRequest{}, nil
	case DeviceEchoResponseType:
		return &DeviceEchoResponse{}, nil
	case DeviceGetGroupType:
		return &DeviceGetGroup{}, nil
	case DeviceGetHostFirmwareType:
		return &DeviceGetHostFirmware{}, nil
	case DeviceGetHostInfoType:
		return &DeviceGetHostInfo{}, nil
	case DeviceGetInfoType:
		return &DeviceGetInfo{}, nil
	case DeviceGetLabelType:
		return &DeviceGetLabel{}, nil
	case DeviceGetLocationType:
		return &DeviceGetLocation{}, nil
	case DeviceGetPowerType:
		return &DeviceGetPower{}, nil
	case DeviceGetServiceType:
		return &DeviceGetService{}, nil
	case DeviceGetVersionType:
		return &DeviceGetVersion{}, nil
	case DeviceGetWifiFirmwareType:
		return &DeviceGetWifiFirmware{}, nil
	case DeviceGetWifiInfoType:
		return &DeviceGetWifiInfo{}, nil
	case DeviceSetGroupType:
		return &DeviceSetGroup{}, nil
	case DeviceSetLabelType:
		return &DeviceSetLabel{}, nil
	case DeviceSetLocationType:
		return &DeviceSetLocation{}, nil
	case DeviceSetPowerType:
		return &DeviceSetPower{}, nil
	case DeviceStateGroupType:
		return &DeviceStateGroup{}, nil
	case DeviceStateHostFirmwareType:
		return &DeviceStateHostFirmware{}, nil
	case DeviceStateHostInfoType:
		return &DeviceStateHostInfo{}, nil
	case DeviceStateInfoType:
		return &DeviceStateInfo{}, nil
	case DeviceStateLabelType:
		return &DeviceStateLabel{}, nil
	case DeviceStateLocationType:
		return &DeviceStateLocation{}, nil
	case DeviceStatePowerType:
		return &DeviceStatePower{}, nil
	case DeviceStateServiceType:
		return &DeviceStateService{}, nil
	case DeviceStateVersionType:
		return &DeviceStateVersion{}, nil
	case DeviceStateWifiFirmwareType:
		return &DeviceStateWifiFirmware{}, nil
	case DeviceStateWifiInfoType:
		return &DeviceStateWifiInfo{}, nil
	case LightGetType:
		return &LightGet{}, nil
	case LightGetInfraredType:
		return &LightGetInfrared{}, nil
	case LightGetPowerType:
		return &LightGetPower{}, nil
	case LightSetColorType:
		return &LightSetColor{}, nil
	case LightSetInfraredType:
		return &LightSetInfrared{}, nil
	case LightSetPowerType:
		return &LightSetPower{}, nil
	case LightSetWaveformType:
		return &LightSetWaveform{}, nil
	case LightSetWaveformOptionalType:
		return &LightSetWaveformOptional{}, nil
	case LightStateType:
		return &LightState{}, nil
	case LightStateInfraredType:
		return &LightStateInfrared{}, nil
	case LightStatePowerType:
		return &LightStatePower{}, nil
	case MultiZoneExtendedGetColorZonesType:
		return &MultiZoneExtendedGetColorZones{}, nil
	case MultiZoneExtendedSetColorZonesType:
		return &MultiZoneExtendedSetColorZones{}, nil
	case MultiZoneExtendedStateMultiZoneType:
		return &MultiZoneExtendedStateMultiZone{}, nil
	case MultiZoneGetColorZonesType:
		return &MultiZoneGetColorZones{}, nil
	case MultiZoneGetEffectType:
		return &MultiZoneGetEffect{}, nil
	case MultiZoneSetColorZonesType:
		return &MultiZoneSetColorZones{}, nil
	case MultiZoneSetEffectType:
		return &MultiZoneSetEffect{}, nil
	case MultiZoneStateEffectType:
		return &MultiZoneStateEffect{}, nil
	case MultiZoneStateMultiZoneType:
		return &MultiZoneStateMultiZone{}, nil
	case MultiZoneStateZoneType:
		return &MultiZoneStateZone{}, nil
	case TileGet64Type:
		return &TileGet64{}, nil
	case TileGetDeviceChainType:
		return &TileGetDeviceChain{}, nil
	case TileGetEffectType:
		return &TileGetEffect{}, nil
	case TileSet64Type:
		return &TileSet64{}, nil
	case TileSetEffectType:
		return &TileSetEffect{}, nil
	case TileSetUserPositionType:
		return &TileSetUserPosition{}, nil
	case TileState64Type:
		return &TileState64{}, nil
	case TileStateDeviceChainType:
		return &TileStateDeviceChain{}, nil
	case TileStateEffectType:
		return &TileStateEffect{}, nil
	}
	return nil, fmt.Errorf("unable to find payload of type %d", payloadType)
}
