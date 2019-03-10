package lan_test

import (
	"testing"

	"github.com/greygore/lifx/lan"
)

func TestPacketStringers(t *testing.T) {
	c1 := lan.LightHsbk{
		Hue:        21845,
		Saturation: 65535,
		Brightness: 65535,
		Kelvin:     3500,
	}
	c2 := lan.LightHsbk{
		Hue:        21845,
		Saturation: 65535,
		Brightness: 0,
		Kelvin:     3500,
	}

	type test struct {
		Name     string
		Payload  lan.Payloader
		Expected string
	}
	tests := []test{
		{
			Name:     "Minimal packet",
			Payload:  &lan.DeviceGetService{},
			Expected: "DeviceGetService",
		},
		{
			Name: "Embedded field with leading comma (first field reserved)",
			Payload: &lan.LightSetColor{
				Color:    c1,
				Duration: 1024,
			},
			Expected: "LightSetColor:Color=LightHsbk:{Hue=21845,Saturation=65535,Brightness=65535,Kelvin=3500},Duration=1024",
		},
		{
			Name: "Embedded field and enum, plus assorted types (bool, ints, floats)",
			Payload: &lan.LightSetWaveformOptional{
				Transient: true,
				Color:     c2,
				Period:    100,
				Cycles:    1.5,
				SkewRatio: 3,
				Waveform:  lan.LightWaveformSine,
			},
			Expected: "LightSetWaveformOptional:Transient=true,Color=LightHsbk:{Hue=21845,Saturation=65535,Brightness=0,Kelvin=3500},Period=100,Cycles=1.500000,SkewRatio=3,Waveform=LIGHT_WAVEFORM_SINE,SetHue=false,SetSaturation=false,SetBrightness=false,SetKelvin=false",
		},
		{
			Name: "Array of embedded fields",
			Payload: &lan.MultiZoneStateMultiZone{
				Count: 3,
				Index: 1,
				Color: [8]lan.LightHsbk{c1, c2, c1, c2, c1, c2, c1, c2},
			},
			Expected: "MultiZoneStateMultiZone:Count=3,Index=1,Color=[LightHsbk:{Hue=21845,Saturation=65535,Brightness=65535,Kelvin=3500} LightHsbk:{Hue=21845,Saturation=65535,Brightness=0,Kelvin=3500} LightHsbk:{Hue=21845,Saturation=65535,Brightness=65535,Kelvin=3500} LightHsbk:{Hue=21845,Saturation=65535,Brightness=0,Kelvin=3500} LightHsbk:{Hue=21845,Saturation=65535,Brightness=65535,Kelvin=3500} LightHsbk:{Hue=21845,Saturation=65535,Brightness=0,Kelvin=3500} LightHsbk:{Hue=21845,Saturation=65535,Brightness=65535,Kelvin=3500} LightHsbk:{Hue=21845,Saturation=65535,Brightness=0,Kelvin=3500}]",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := test.Payload.String()
			if actual != test.Expected {
				t.Errorf("Stringer mismatch:\nExpected:\t'%s'\nActual:\t'%s'", test.Expected, actual)
			}
		})
	}
}
