package main

import (
	"log"
	"regexp"
	"sort"
)

type enum struct {
	Name   string `yaml:"-"`
	Type   string
	Values []struct {
		Name  string
		Value int
	}
}

const enumTemplate = `
{{- range . -}}
// {{pascal .Name}} is generated from protocol enums
type {{pascal .Name}} {{.Type}}

// {{pascal .Name}} values
const (
	{{- range .Values}}
	{{pascal .Name}} = {{.Value}}
	{{- end}}
)

func (e {{pascal .Name}}) String() string {
	switch e {
	{{- range .Values}}
	case {{pascal .Name}}:
		return "{{.Name}}"
	{{- end}}
	}
	return fmt.Sprintf("{{pascal .Name}}(%d)", e)
}

{{end}}
`

var (
	enumRegex = regexp.MustCompile(`^(\[[1-9][0-9]*\])?<([A-Za-z0-9]+)>$`)
)

func genEnums(data map[string]enum) {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Convert map to sorted slice and remove reserved values
	var enums []enum
	for _, key := range keys {
		e := enum{Name: key, Type: data[key].Type}
		for i := range data[key].Values {
			if data[key].Values[i].Name == "reserved" {
				continue
			}
			e.Values = append(e.Values,
				struct {
					Name  string
					Value int
				}{
					data[key].Values[i].Name,
					data[key].Values[i].Value,
				},
			)
		}
		enums = append(enums, e)
	}

	src := NewGenerator("enums", "", enumTemplate)
	src.Printf("import \"fmt\"\n\n")

	if err := src.WriteFile(enums); err != nil {
		log.Fatalf("Unable to write enums.go:\n\t%s", err)
	}
}
