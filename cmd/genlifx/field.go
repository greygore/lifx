package main

import (
	"log"
	"sort"
)

type field struct {
	Name      string `yaml:"-"`
	SizeBytes int    `yaml:"size_bytes"`
	Fields    []struct {
		Name      string
		Type      string
		SizeBytes int `yaml:"size_bytes"`
	}
}

const fieldTemplate = `
{{- range . -}}
type {{pascal .Name}} struct {
	{{- range .Fields}}
	{{.Name}} {{.Type}}
	{{- end}}
}

{{end}}
`

func genFields(data map[string]field) {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Convert map to sorted slice and remove reserved values
	var fields []field
	for _, key := range keys {
		f := field{Name: key, SizeBytes: data[key].SizeBytes}
		for i := range data[key].Fields {
			if data[key].Fields[i].Type == "reserved" {
				continue
			}
			fieldType := data[key].Fields[i].Type
			enumType := enumRegex.FindStringSubmatch(fieldType)
			if len(enumType) > 2 {
				fieldType = enumType[1] + enumType[2]
			}
			f.Fields = append(f.Fields,
				struct {
					Name      string
					Type      string
					SizeBytes int `yaml:"size_bytes"`
				}{
					data[key].Fields[i].Name,
					fieldType,
					data[key].Fields[i].SizeBytes,
				},
			)
		}
		fields = append(fields, f)
	}

	src := NewGenerator("fields", "", fieldTemplate)

	if err := src.WriteFile(fields); err != nil {
		log.Fatalf("Unable to write fields.go:\n\t%s", err)
	}
}
