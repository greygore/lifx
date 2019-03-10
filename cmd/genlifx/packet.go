package main

import (
	"log"
	"sort"
)

type packet struct {
	Name       string `yaml:"-"`
	Kind       string `yaml:"-"`
	PacketType int    `yaml:"pkt_type"`
	SizeBytes  int    `yaml:"size_bytes"`
	Fields     []struct {
		Name      string
		Type      string
		SizeBytes int `yaml:"size_bytes"`
	}
}

const payloadTemplate = `
func PayloadByType(payloadType uint16) (Payloader, error) {
	switch payloadType {
	{{- range .}}{{range .}}
	case {{.Kind}}Type:
		return &{{.Kind}}{}, nil
	{{- end}}{{end}}
	}
	return nil, fmt.Errorf("unable to find payload of type %d", payloadType)
}
`

const packetTemplate = `// message types
const (
	{{- range .}}
	{{.Kind}}Type = {{.PacketType}}
	{{- end}}
)

{{range . -}}
/////////////////////////////////////////////////////////////////////////////

type {{.Kind}} struct{
{{- if .Fields}}
	{{- range $i, $k := .Fields}}
	{{if .Name}}{{.Name}} {{.Type}}{{else}}Reserved{{$i}}[{{.SizeBytes}}]byte{{end}}
	{{- end}}
{{end -}}
}

func (m {{.Kind}}) Size() uint16 {
	return {{.SizeBytes}}
}

func (m {{.Kind}}) Type() uint16 {
	return {{.Kind}}Type
}

func (m {{.Kind}}) MarshalBinary() ([]byte, error) {
	{{if .Fields -}}
	b := &bytes.Buffer{}

	data := []interface{}{
		{{- range $i, $k := .Fields}}
		{{if .Name}}m.{{.Name}},{{else}}m.Reserved{{$i}},{{end}}
		{{- end}}
	}
	for _, d := range data {
		if err := binary.Write(b, endian, d); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
	{{- else -}}
	return []byte{}, nil
	{{- end}}
}

func (m *{{.Kind}}) UnmarshalBinary(data []byte) error {
	{{if .Fields -}}
	if int(m.Size()) != len(data) {
		return fmt.Errorf("expected %d bytes, got %d", m.Size(), len(data))
	}

	b := bytes.NewBuffer(data)
	vars := []interface{}{
		{{- range $i, $k := .Fields}}
		{{if .Name}}&m.{{.Name}},{{else}}&m.Reserved{{$i}},{{end}}
		{{- end}}
	}
	for _, v := range vars {
		if err := binary.Read(b, endian, v); err != nil {
			return fmt.Errorf("unable to read packet: %s", err)
		}
	}
	{{- else -}}
	if len(data) > 0 {
		return fmt.Errorf("expected empty packet, got %d bytes", len(data))
	}
	{{- end}}

	return nil
}

func (m {{.Kind}}) String() string {
	{{if .Fields -}}
	var s strings.Builder
	s.WriteString("{{.Kind}}:")

	{{$comma := false}}
	{{- range .Fields}}
	{{if .Name -}}
		s.WriteString(fmt.Sprintf("{{if $comma}},{{end}}{{.Name}}={{formatchar .Type}}", m.{{.Name}}))
		{{- $comma = true -}}
	{{- end}}
	{{- end}}

	return s.String()
	{{- else -}}
	return "{{.Kind}}"
	{{- end}}
}

{{end}}
`

func genPackets(data map[string]map[string]packet) {
	generators := map[string]*Generator{}
	keys := map[string][]string{}
	for name := range data {
		generators[name] = NewGenerator("packets", name, packetTemplate)
		generators[name].Printf("import (\n\t\"bytes\"\n\t\"encoding/binary\"\n\t\"fmt\"\n\t\"strings\"\n)\n\n")
		for kind := range data[name] {
			keys[name] = append(keys[name], kind)
		}
	}
	for name := range keys {
		sort.Strings(keys[name])
	}

	// Convert map to sorted slice
	var largestPacketSize int
	packets := map[string][]packet{}
	for name := range keys {
		for _, kind := range keys[name] {
			if data[name][kind].SizeBytes > largestPacketSize {
				largestPacketSize = data[name][kind].SizeBytes
			}
			p := packet{
				Name:       name,
				Kind:       kind,
				PacketType: data[name][kind].PacketType,
				SizeBytes:  data[name][kind].SizeBytes,
			}

			for i := range data[name][kind].Fields {
				fieldType := data[name][kind].Fields[i].Type
				enumType := enumRegex.FindStringSubmatch(fieldType)
				if len(enumType) > 2 {
					fieldType = enumType[1] + enumType[2]
				}
				p.Fields = append(p.Fields,
					struct {
						Name      string
						Type      string
						SizeBytes int `yaml:"size_bytes"`
					}{
						data[name][kind].Fields[i].Name,
						fieldType,
						data[name][kind].Fields[i].SizeBytes,
					},
				)
			}

			packets[name] = append(packets[name], p)
		}
	}

	for name := range keys {
		if err := generators[name].WriteFile(packets[name]); err != nil {
			log.Fatalf("Unable to write packets_%s.go:\n\t%s", name, err)
		}
	}

	g := NewGenerator("packets", "", payloadTemplate)
	g.Printf("import \"fmt\"\n\n")
	g.Printf("const largestPacketSize = %d\n\n", largestPacketSize)
	if err := g.WriteFile(packets); err != nil {
		log.Fatalf("Unable to write packets.go:\n\t%s", err)
	}
}
